package main

import (
	"fmt"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/spf13/cobra"
)

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect",
	RunE: func(cmd *cobra.Command, args []string) error {
		repos := repository.NewRepos()
		feedSrv := service.NewFeedSevice(repos)
		programSrv := service.NewProgramService(repos)

		feeds := feedSrv.List()
		for _, feed := range feeds {
			fmt.Printf("found: %s\n", feed.Url)
			realfeed, err := feedSrv.Refetch(feed.Id)
			if err != nil {
				return err
			}
	
			for _, realfeeditem := range realfeed.Items {
				id := programSrv.Create(service.Program{
					Title:     realfeeditem.Title,
					Content:   realfeeditem.Content,
					Converted: false,
				})
				fmt.Printf("program %s created. title: %s\n", id, realfeeditem.Title)
				if err := programSrv.Convert(id); err != nil {
					return err
				}
				fmt.Printf("program %s converted.\n", id)
			}
		}

		return nil
	},
}
