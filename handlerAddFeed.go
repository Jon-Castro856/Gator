package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *config.State, c config.Command, user database.User) error {
	args := c.Args
	ctx := context.Background()

	if len(args) != 2 {
		return fmt.Errorf("not enough arguments in the command: require name and URL")
	}
	name := args[0]
	feedURL := args[1]

	feed, err := s.DB.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedName:  name,
		Url:       feedURL,
		UserID:    user.ID,
	})

	if err != nil {
		return err
	}

	_, err = s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("succesfully created feed")
	fmt.Println(feed)
	return nil
}
