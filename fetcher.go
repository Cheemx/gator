package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Cheemx/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Name of the feed: %s\n", feed.Name)
	fmt.Println()
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}
	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, item := range fetchedFeed.Channel.Item {
		layout := "Mon, 02 Jan 2006 15:04:05 MST"
		pubTime, err := time.Parse(layout, item.PubDate)
		if err != nil {
			return err
		}
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: pubTime, Valid: true},
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			return err
		}
		fmt.Printf("Post inserted into DB: %+v\n", post)
	}
	fmt.Println("")
	return nil
}
