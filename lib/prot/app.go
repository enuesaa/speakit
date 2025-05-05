package prot

import (
	"fmt"
	"reflect"
)

type Record struct {
	Text  string
	Voice []byte
	Meta  map[string]string
}

type App struct {
	logger       Logger
	notify       NotifyBehavior
	generator    Generator
	skippers     []Skipper
	transformers []Transformer
	controllers  []Controller
	speaker      Speaker
}

func (a *App) Generate(generator Generator) {
	a.generator = generator
}

func (a *App) Skipper(skipper Skipper) {
	a.skippers = append(a.skippers, skipper)
}

func (a *App) Transform(transformer Transformer) {
	a.transformers = append(a.transformers, transformer)
}

func (a *App) Controller(controller Controller) {
	a.controllers = append(a.controllers, controller)
}

func (a *App) Speak(speaker Speaker) {
	a.speaker = speaker
}

func (a *App) Run() error {
	a.logger = Logger{}
	a.notify = newNotifyBehavior(a.speaker)

	if err := a.startUp(); err != nil {
		return err
	}
	defer a.close()

	var occured error
	for {
		record, err := a.generator.Generate()
		if err != nil {
			occured = err
			break
		}
		if a.shouldSkip(record) {
			continue
		}
		if err := a.transformRecord(&record); err != nil {
			occured = err
			break
		}
		a.notify.waitIfNeed()

		if err := a.speaker.Speak(record); err != nil {
			occured = err
			break
		}
	}
	return occured
}

func (a *App) shouldSkip(record Record) bool {
	for _, s := range a.skippers {
		if s.ShouldSkip(record) {
			return true
		}
	}
	return false
}

func (a *App) transformRecord(record *Record) error {
	for _, t := range a.transformers {
		if err := t.Transform(record); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) startUp() error {
	ilog := func(i any) Logger {
		return a.logger.Use(i)
	}

	if err := a.generator.StartUp(ilog(a.generator)); err != nil {
		return err
	}
	for _, s := range a.skippers {
		if err := s.StartUp(ilog(s)); err != nil {
			return err
		}
	}
	for _, t := range a.transformers {
		if err := t.StartUp(ilog(t)); err != nil {
			return err
		}
	}
	for _, c := range a.controllers {
		if err := c.StartUp(ilog(c)); err != nil {
			return err
		}
	}
	if err := a.speaker.StartUp(ilog(a.speaker)); err != nil {
		return err
	}
	return nil
}

func (a *App) listfns() []any {
	fns := []any{a.generator, a.speaker}
	for _, t := range a.transformers {
		fns = append(fns, t)
	}
	for _, c := range a.controllers {
		fns = append(fns, c)
	}
	for _, s := range a.skippers {
		fns = append(fns, s)
	}
	return fns
}

func (a *App) callInject() error {
	for _, fn := range a.listfns() {
		fn := reflect.ValueOf(fn).MethodByName("Inject")
		if !fn.IsValid() {
			return nil
		}
		sig := fn.Type()
		args := make([]reflect.Value, 0)
	
		for i := range sig.NumIn() {
			switch sig.In(i) {
			case reflect.TypeOf(a.logger):
				logger := a.logger.Use(a.generator)
				args = append(args, reflect.ValueOf(logger))
			default:
				return fmt.Errorf("unsupported: %v", sig)
			}
		}
		fn.Call(args)	
	}
	return nil
}

func (a *App) close() {
	if err := a.generator.Close(); err != nil {
		a.logger.LogE(err)
	}
	for _, s := range a.skippers {
		if err := s.Close(); err != nil {
			a.logger.LogE(err)
		}
	}
	for _, t := range a.transformers {
		if err := t.Close(); err != nil {
			a.logger.LogE(err)
		}
	}
	for _, c := range a.controllers {
		if err := c.Close(); err != nil {
			a.logger.LogE(err)
		}
	}
	if err := a.speaker.Close(); err != nil {
		a.logger.LogE(err)
	}
}
