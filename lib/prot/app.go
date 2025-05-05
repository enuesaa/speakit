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
	log          *LogBehavior
	notify       *NotifyBehavior
	generator    Generator
	skippers     []Skipper
	transformers []Transformer
	controllers  []Controller
	speaker      Speaker
}

func (a *App) Generate(generator Generator) {
	a.generator = generator
}

func (a *App) Skip(skipper Skipper) {
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
	a.log = newLogBehavior()
	a.notify = newNotifyBehavior(a.speaker)

	if err := a.callInject(); err != nil {
		return err
	}
	if err := a.callStartUp(); err != nil {
		return err
	}
	defer a.callClose()

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

func (a *App) listCallfns() []Callfn {
	fns := []Callfn{a.generator}
	for _, t := range a.transformers {
		fns = append(fns, t)
	}
	for _, c := range a.controllers {
		fns = append(fns, c)
	}
	for _, s := range a.skippers {
		fns = append(fns, s)
	}
	fns = append(fns, a.speaker)
	return fns
}

func (a *App) callInject() error {
	behaviors := make(map[reflect.Type]reflect.Value)
	behaviors[reflect.TypeOf(a.notify)] = reflect.ValueOf(a.notify)
	behaviors[reflect.TypeOf(&PwBehavior{})] = reflect.ValueOf(&PwBehavior{})

	for _, fn := range a.listCallfns() {
		fn := reflect.ValueOf(fn).MethodByName("Inject")
		if !fn.IsValid() {
			return nil
		}
		behaviors[reflect.TypeOf(a.log)] = reflect.ValueOf(a.log.Use(fn))
		sig := fn.Type()
		var args []reflect.Value

		for i := range sig.NumIn() {
			behavior, ok := behaviors[sig.In(i)]
			if !ok {
				return fmt.Errorf("unsupported: %v", sig)
			}
			args = append(args, behavior)
		}
		fn.Call(args)
	}
	return nil
}

func (a *App) callStartUp() error {
	for _, callfn := range a.listCallfns() {
		if err := callfn.StartUp(); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) callClose() {
	for _, callfn := range a.listCallfns() {
		if err := callfn.Close(); err != nil {
			a.log.LogE(err)
		}
	}
}
