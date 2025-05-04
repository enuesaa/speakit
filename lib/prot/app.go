package prot

type Record struct {
	Text string
	Voice []byte
	Meta map[string]string
}

type App struct {
	generator Generator
	transformers []Transformer
	controllers []Controller
	speaker Speaker
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

func (a *App) Prev() error {
	return nil
}

func (a *App) Next() error {
	return nil
}

func (a *App) Stop() error {
	return a.speaker.Stop()
}

func (a *App) Run() error {
	if err := a.Start(); err != nil {
		return err
	}

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
		if err := a.speaker.Speak(record); err != nil {
			occured = err
			break
		}
	}
	return occured
}

func (a *App) Start() error {
	if err := a.generator.StartUp(a); err != nil {
		return err
	}
	for _, t := range a.transformers {
		if err := t.StartUp(a); err != nil {
			return err
		}
	}
	for _, c := range a.controllers {
		if err := c.StartUp(a); err != nil {
			return err
		}
	}
	if err := a.speaker.StartUp(a); err != nil {
		return err
	}
	return nil
}

func (a *App) transformRecord(record *Record) error {
	for _, t := range a.transformers {
		if err := t.Transform(record); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) Logger(name string) Logger {
	return Logger{name: name}
}
