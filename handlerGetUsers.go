package main

import (
	"context"
	"fmt"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func handlerGetUsers(s *config.State, c config.Command) error {
	ctx := context.Background()

	users, err := s.DB.GetUsers(ctx)
	if err != nil {
		fmt.Printf("error getting users: %v\n", err)
	}

	for _, user := range users {
		if user == s.Config.Username {
			fmt.Println(user + " (current)")
		} else {
			fmt.Println(user)
		}
	}
	return nil
}
