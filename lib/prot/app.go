package prot

import "time"

type Record struct {
	Text  string
	Voice []byte
	Meta  map[string]string
}

type App struct {
	wait         bool
	logger       Logger
	generator    Generator
	transformers []Transformer
	controllers  []Controller
	speaker      Speaker
}

func (a *App) Generate(generator Generator) {
	a.generator = generator
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

func (a *App) Next() error {
	a.wait = false

	return a.speaker.CancelWait()
}

func (a *App) Stop() error {
	a.wait = true

	return a.speaker.CancelWait()
}

func (a *App) Run() error {
	a.logger = Logger{}

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
		if err := a.transformRecord(&record); err != nil {
			occured = err
			break
		}
		a.waitIfNeed()

		if err := a.speaker.Speak(record); err != nil {
			occured = err
			break
		}
	}
	return occured
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
	for _, t := range a.transformers {
		if err := t.StartUp(ilog(t)); err != nil {
			return err
		}
	}
	for _, c := range a.controllers {
		if err := c.StartUp(ilog(c), a); err != nil {
			return err
		}
	}
	if err := a.speaker.StartUp(ilog(a.speaker)); err != nil {
		return err
	}
	return nil
}

func (a *App) close() {
	if err := a.generator.Close(); err != nil {
		a.logger.LogE(err)
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

func (a *App) waitIfNeed() {
	if a.wait {
		for {
			if !a.wait {
				break
			}
			time.Sleep(3 * time.Second)
		}
	} else {
		// speaker
		for {
			if a.speaker.IsStopped() {
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}
