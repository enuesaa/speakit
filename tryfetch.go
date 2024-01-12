package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
)

type SpeakitConfig struct {
	Url string `json:"url"`
	OpenAiApiKey string `json:"openAiApiKey"`
	PromptTemplate string `json:"promptTemplate"`
	Results []SpeakitResult `json:"speakitResults"`
}

type SpeakitResult struct {
	Output string `json:"output"`
	Title string `json:"title"`
	Url string `json:"url"`
}

func TryFetch() {
	env := repository.Env{}
	repos := repository.NewRepos(env)

	if !repos.Fs.IsExist("speakit.json") {
		b, err := json.MarshalIndent(SpeakitConfig{
			Results: make([]SpeakitResult, 0),
		}, "", "  ")
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		if err := repos.Fs.Create("speakit.json", b); err != nil {
			log.Fatalf("Error: %s", err.Error())			
		}
	}

	configbytes, err := repos.Fs.Read("speakit.json")
	if err != nil {
		log.Fatalf("Error: failed to open config file. %s", err.Error())
	}
	var config SpeakitConfig
	if err := json.Unmarshal(configbytes, &config); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	fmt.Printf("url: %s\n", config.Url)

	feedSrv := service.NewFeedSevice(repos)
	feeds, err := feedSrv.TryFetch(config.Url)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	for _, feeditem := range feeds.Items {
		fmt.Printf("found: %s\n", feeditem.Title)
	}
}
