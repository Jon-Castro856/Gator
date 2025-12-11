package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *config.State, cmd config.Command) error {
	args := cmd.Args
	ctx := context.Background()

	if len(args) != 1 {
		return fmt.Errorf("no arguments in the command")
	}
	name := args[0]

	_, err := s.DB.GetUser(ctx, name)
	if err == nil {
		fmt.Println("User already exists")
		os.Exit(1)
	}

	user, err := s.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserName:  args[0],
	})

	if err != nil {
		fmt.Printf("error with creating new user: %v\n", err)
		os.Exit(1)
	}

	s.Config.SetUser(user.UserName)
	fmt.Println("New user registered")
	fmt.Printf("User name: %v\nUser ID: %v\nCreated at: %v\n", user.UserName, user.ID, user.CreatedAt)

	return nil
}
