package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func handlerBrowse(s *config.State, c config.Command) error {
	args := c.Args
	ctx := context.Background()
	var limit int32

	if len(args) > 0 {
		limit = convertToInt(args[0])
	} else {
		limit = 2
	}

	fmt.Printf("Showing %v posts\n", limit)
	postList, err := s.DB.GetPosts(ctx, int32(limit))
	if err != nil {
		return err
	}

	for _, post := range postList {
		fmt.Println(post.Title)
		fmt.Println(post.Url)
	}
	return nil
}

func convertToInt(s string) int32 {
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 2
	}
	return int32(i64)
}
