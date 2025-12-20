package main

import (
	"context"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
)

func middlewareLoggedIn(handler func(s *config.State, c config.Command, user database.User) error) func(*config.State, config.Command) error {
	return func(s *config.State, c config.Command) error {
		user, err := s.DB.GetUserInfo(context.Background(), s.Config.Username)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}
