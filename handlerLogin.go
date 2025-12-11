package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func handlerLogin(s *config.State, cmd config.Command) error {
	args := cmd.Args
	ctx := context.Background()

	if len(args) != 1 {
		return fmt.Errorf("no arguments in the command")
	}
	name := args[0]

	_, err := s.DB.GetUser(ctx, name)

	if err != nil {
		fmt.Printf("error logging in: %v\n", err)
		os.Exit(1)
	}

	s.Config.SetUser(name)
	fmt.Println("user has been set")
	return nil
}
