package prot

import (
	"fmt"
	"time"
)

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
	return nil
}

func (a *App) Run() error {
	if err := a.Start(); err != nil {
		return err
	}
	records, err := a.generator.Generate()
	if err != nil {
		return err
	}
	fmt.Printf("len: %d\n", len(records))

	for i, record := range records {
		fmt.Printf("record: %d\n", i)
		if err := a.transformRecord(&record); err != nil {
			return err
		}
		if err := a.speaker.Speak(record); err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

func (a *App) Start() error {
	for _, t := range a.transformers {
		if err := t.StartUp(); err != nil {
			return err
		}
	}
	for _, c := range a.controllers {
		if err := c.StartUp(a); err != nil {
			return err
		}
	}
	if err := a.speaker.StartUp(); err != nil {
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

// func (a *App) Next() {}
// func (a *App) Stop() {}
// func (a *App) Close() {}
