package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *config.State, c config.Command, user database.User) error {
	args := c.Args
	ctx := context.Background()

	if len(args) != 1 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	url := args[0]

	feedInfo, err := s.DB.GetFeedNameAndID(ctx, url)
	if err != nil {
		fmt.Printf("error retreiving feed name: %v\n", err)
	}

	feedFollow, err := s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feedInfo.ID,
	})
	if err != nil {
		fmt.Printf("error in created feed_follow entry: %v\n", err)
	}

	fmt.Println("following new feed:")
	fmt.Printf("Feed name: %v\nFollower Name: %v\n", feedFollow.FeedName, feedFollow.UserName)
	return nil
}
