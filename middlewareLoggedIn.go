package main

import (
	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
)

func middlewareLoggedIn(handler func(s *config.State, c config.Command, user database.User) error) func(*config.State, config.Command) error {
	userName := s.Config.UserName
}
