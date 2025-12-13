package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func handlerAgg(s *config.State, c config.Command) error {
	url := "https://www.wagslane.dev/index.xml"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	feed, err := fetchFeed(ctx, url)
	if err != nil {
		fmt.Printf("error retreiving RSS feed: %v", err)
		os.Exit(1)
	}
	fmt.Println(feed)
	return nil
}
