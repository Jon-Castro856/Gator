package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Jon-Castro856/Gator/internal/rss"
)

func fetchFeed(ctx context.Context, feedURL string) (*rss.RSSFeed, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		fmt.Printf("error retrieving web data: %v\n", err)
	}
	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("more errors, god: %v\n", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading data from request: %v\n", err)
	}

	var feed rss.RSSFeed
	if err := xml.Unmarshal(data, &feed); err != nil {
		fmt.Printf("error unmarshaling data: %v\n", err)
		os.Exit(1)
	}

	return &feed, nil
}
