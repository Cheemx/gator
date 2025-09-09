package main

import (
	"context"
	"fmt"
	"log"
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
		log.Fatal(err)
	}
	for _, item := range fetchedFeed.Channel.Item {
		fmt.Println(item.Title)
	}
	fmt.Println("")
	return nil
}
