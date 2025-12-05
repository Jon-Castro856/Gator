package main

import (
	"fmt"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func handlerLogin(s *config.State, cmd config.Command) error {
	args := cmd.Args
	if len(args) != 1 {
		return fmt.Errorf("no arguments in the command")
	}

	s.Config.SetUser(args[0])
	fmt.Println("user has been set")
	return nil
}
