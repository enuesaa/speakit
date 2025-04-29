package prot

import "fmt"

type Record struct {
	Text string
	Voice []byte
	Meta map[string]string
}

type App struct {
	generator Generator
	transformers []Transformer
	speaker Speaker
}

func (a *App) Transform(transformer Transformer) {
	a.transformers = append(a.transformers, transformer)
}

func (a *App) Speak(speaker Speaker) error {
	a.speaker = speaker
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
		if err := a.speaker.Next(record); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) Start() error {
	for _, t := range a.transformers {
		if err := t.Start(); err != nil {
			return err
		}
	}
	if err := a.speaker.Start(); err != nil {
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
