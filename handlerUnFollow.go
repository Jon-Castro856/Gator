package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
)

func handlerUnFollow(s *config.State, c config.Command, user database.User) error {
	args := c.Args
	if len(args) != 1 {
		fmt.Println("no url included in argument")
		os.Exit(1)
	}
	url := args[0]
	ctx := context.Background()

	err := s.DB.UnfollowFeed(ctx, url)
	if err != nil {
		return err
	}
	fmt.Printf("no longer following feed at %v\n", url)
	return nil
}
