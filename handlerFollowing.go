package main

import (
	"context"
	"fmt"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func handlerFollowing(s *config.State, c config.Command) error {
	ctx := context.Background()

	userName := s.Config.Username

	userId, err := s.DB.GetUserID(ctx, userName)
	if err != nil {
		fmt.Printf("error retreiving id: %v\n", err)
	}

	feedFollows, err := s.DB.GetFeedFollowsForUser(ctx, userId)
	if err != nil {
		fmt.Printf("error retrieving feed follows: %v\n", err)
	}

	fmt.Printf("Displaying feeds for %v\n", userName)
	fmt.Println(feedFollows)

	return nil
}
