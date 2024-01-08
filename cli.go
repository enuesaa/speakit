package main

import (
	"fmt"
	"log"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
)

func TryFetch(url string) {
	env := repository.Env{}
	repos := repository.NewRepos(env)

	feedSrv := service.NewFeedSevice(repos)
	feeds, err := feedSrv.TryFetch(url)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	for _, feeditem := range feeds.Items {
		fmt.Printf("found: %s\n", feeditem.Title)
	}
}
