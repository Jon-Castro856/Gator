package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *config.State, c config.Command) error {
	args := c.Args
	ctx := context.Background()

	if len(args) != 2 {
		return fmt.Errorf("not enough arguments in the command: require name and URL")
	}
	name := args[0]
	feedURL := args[1]

	id, err := s.DB.GetUserID(context.Background(), s.Config.Username)
	if err != nil {
		fmt.Printf("error getting user name: %v\n", err)
		os.Exit(1)
	}

	feed, err := s.DB.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedName:  name,
		Url:       feedURL,
		UserID:    id,
	})

	if err != nil {
		fmt.Printf("error creating feed: %v\n", err)
	}

	_, err = s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: id,
		FeedID: feed.ID,
	})
	if err != nil {
		fmt.Printf("error creating feed follow entry: %v\n", err)
	}
	fmt.Println("succesfully created feed")
	fmt.Println(feed)
	return nil
}
