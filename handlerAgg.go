package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func handlerAgg(s *config.State, c config.Command) error {
	args := c.Args
	if len(args) != 2 {
		fmt.Println("incorrect arguments")
		os.Exit(1)
	}

	if args[0] != "tbr" {
		fmt.Println("use 'tbr' to indicate time between requests")
		os.Exit(1)
	}

	timer, err := time.ParseDuration(args[1])
	if err != nil {
		return err
	}

	fmt.Printf("Now scraping feeds every %v\n", timer)
	ticker := time.NewTicker(timer)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *config.State) error {
	ctx := context.Background()
	nextFeed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	err = s.DB.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return err
	}

	webCTX, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	fetched, err := fetchFeed(webCTX, nextFeed.Url)
	if err != nil {
		return err
	}

	for _, feed := range fetched.Channel.Item {
		err := createPost(s, feed, nextFeed.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
