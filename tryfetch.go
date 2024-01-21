package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type SpeakitConfig struct {
	Url            string          `json:"url"`
	OpenAiApiKey   string          `json:"openAiApiKey"`
	PromptTemplate string          `json:"promptTemplate"`
	Results        []SpeakitResult `json:"speakitResults"`
}

type SpeakitResult struct {
	Id     string `json:"id"`
	Output string `json:"output"`
	Title  string `json:"title"`
	Url    string `json:"url"`
}

var tryFetchCmd = &cobra.Command{
	Use:   "try-fetch",
	Short: "try to fetch rss feed",
	Run: func(cmd *cobra.Command, args []string) {
		env := repository.Env{}
		repos := repository.NewRepos(env)
		repos.Voicevox.SetBaseUrl("http://localhost:50021")
	
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
	
		if err := repos.Fs.Remove("data"); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		if err := repos.Fs.CreateDir("data"); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	
		config.Results = make([]SpeakitResult, 0)
		for i, feeditem := range feeds.Items {
			if i > 2 {
				break
			}
			fmt.Printf("found: %s\n", feeditem.Title)
			id := uuid.NewString()
			config.Results = append(config.Results, SpeakitResult{
				Id:     id,
				Title:  feeditem.Title,
				Url:    feeditem.Link,
				Output: "",
			})
			fmt.Printf("try audioquery\n")
			audioquery, err := repos.Voicevox.AudioQuery(feeditem.Title + "。 " + feeditem.Description)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue
			}
			fmt.Printf("try sythesis\n")
			converted, err := repos.Voicevox.Synthesis(audioquery)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue
			}
			filename := fmt.Sprintf("data/%s.wav", id)
			if err := repos.Fs.Create(filename, []byte(converted)); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
		}
	
		if err := createSpeakitJson(config); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	},
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
