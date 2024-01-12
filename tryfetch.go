package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/google/uuid"
)

type SpeakitConfig struct {
	Url string `json:"url"`
	OpenAiApiKey string `json:"openAiApiKey"`
	PromptTemplate string `json:"promptTemplate"`
	Results []SpeakitResult `json:"speakitResults"`
}

type SpeakitResult struct {
	Id string `json:"id"`
	Output string `json:"output"`
	Title string `json:"title"`
	Url string `json:"url"`
}

func TryFetch() {
	env := repository.Env{}
	repos := repository.NewRepos(env)

	if !repos.Fs.IsExist("speakit.json") {
		newconf := SpeakitConfig{
			Results: make([]SpeakitResult, 0),
		}
		if err := createSpeakitJson(newconf); err != nil {
			log.Fatalf("Error: %s", err.Error())			
		}
	}

	config, err := readSpeakitJson()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	fmt.Printf("url: %s\n", config.Url)

	feedSrv := service.NewFeedSevice(repos)
	feeds, err := feedSrv.TryFetch(config.Url)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	config.Results = make([]SpeakitResult, 0)
	for i, feeditem := range feeds.Items {
		fmt.Printf("found: %s\n", feeditem.Title)
		config.Results = append(config.Results, SpeakitResult{
			Id: uuid.NewString(),
			Title: feeditem.Title,
			Url: feeditem.Link,
			Output: "",
		})
		if i > 4 {
			break
		}
	}
	if err := createSpeakitJson(config); err != nil {
		log.Fatalf("Error: %s", err.Error())			
	}
}

func createSpeakitJson(config SpeakitConfig) error {
	env := repository.Env{}
	repos := repository.NewRepos(env)

	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	if err := repos.Fs.Create("speakit.json", b); err != nil {
		return err		
	}
	return nil
}

func readSpeakitJson() (SpeakitConfig, error) {
	env := repository.Env{}
	repos := repository.NewRepos(env)

	configbytes, err := repos.Fs.Read("speakit.json")
	if err != nil {
		return SpeakitConfig{}, err
	}

	var config SpeakitConfig
	if err := json.Unmarshal(configbytes, &config); err != nil {
		return SpeakitConfig{}, err
	}
	return config, nil
}
