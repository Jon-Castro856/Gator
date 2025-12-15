package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func handlerFeeds(s *config.State, c config.Command) error {
	ctx := context.Background()

	feeds, err := s.DB.GetFeeds(ctx)
	if err != nil {
		fmt.Println("error retreiving feeds: %v\n", err)
		os.Exit(1)
	}

	for _, feed := range feeds {
		user_name, err := s.DB.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			fmt.Println("error getting user id, aborting")
			os.Exit(1)
		}
		fmt.Printf("Feed Name: %s\nURL: %s\nUser Name: %s\n", feed.FeedName, feed.Url, user_name)
	}
	return nil
}
