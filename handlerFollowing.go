package main

import (
	"context"
	"fmt"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
)

func handlerFollowing(s *config.State, c config.Command, user database.User) error {
	ctx := context.Background()

	feedFollows, err := s.DB.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		fmt.Printf("error retrieving feed follows: %v\n", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("no feeds currently being followed")
	} else {
		fmt.Printf("Displaying feeds for %v\n", user.UserName)
		for _, feed := range feedFollows {
			fmt.Printf("%v\n", feed.FeedName)
		}
	}
	return nil
}
