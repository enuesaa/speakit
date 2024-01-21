package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/spf13/cobra"
)

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect",
	Run: func(cmd *cobra.Command, args []string) {
		env := repository.Env{
			MINIO_BUCKET: os.Getenv("MINIO_BUCKET"),
			MINIO_HOST:   os.Getenv("MINIO_HOST"),
			REDIS_HOST:   os.Getenv("REDIS_HOST"),
		}
		fmt.Printf("%s", os.Getenv("REDIS_HOST"))

		repos := repository.NewRepos(env)
		repos.Voicevox.SetBaseUrl("http://localhost:50021")
		feedSrv := service.NewFeedSevice(repos)
		programSrv := service.NewProgramService(repos)

		feeds := feedSrv.List()
		for _, feed := range feeds {
			fmt.Printf("found: %s\n", feed.Url)
			realfeed, err := feedSrv.Refetch(feed.Id)
			if err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
	
			for _, realfeeditem := range realfeed.Items {
				id := programSrv.Create(service.Program{
					Title:     realfeeditem.Title,
					Content:   realfeeditem.Content,
					Converted: false,
				})
				fmt.Printf("program %s created. title: %s\n", id, realfeeditem.Title)
				if err := programSrv.Convert(id); err != nil {
					log.Fatalf("Error: %s", err.Error())
				}
				fmt.Printf("program %s converted.\n", id)
			}
		}
	},
}
