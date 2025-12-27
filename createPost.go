package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
	"github.com/Jon-Castro856/Gator/internal/rss"
	"github.com/google/uuid"
)

func createPost(s *config.State, rss rss.RSSItem, feedid uuid.UUID) error {
	ctx := context.Background()
	descString := makenullString(rss.Description)
	pubDate := makeNullTime(rss.Pubdate)

	params := database.CreatePostParams{
		ID:          uuid.New(),
		Title:       rss.Title,
		Url:         rss.Link,
		Description: descString,
		PublishedAt: pubDate,
		FeedID:      feedid,
	}
	post, err := s.DB.CreatePost(ctx, params)
	if err != nil {
		return err
	}

	fmt.Println("new post created")
	fmt.Println(post)

	return nil
}

func makenullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func makeNullTime(s string) sql.NullTime {
	if s == "" {
		return sql.NullTime{Valid: false}
	}
	t, err := time.Parse("1993-06-15", s)
	if err != nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}
