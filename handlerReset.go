package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func handlerReset(s *config.State, c config.Command) error {
	ctx := context.Background()
	err := s.DB.ResetTable(ctx)

	if err != nil {
		fmt.Printf("error deleting the database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("successly deleted database contents")
	return nil
}
