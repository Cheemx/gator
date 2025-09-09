package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Cheemx/gator/internal/config"
	"github.com/Cheemx/gator/internal/database"
	"github.com/google/uuid"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		log.Fatal("username is required")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[1])
	if err != nil {
		log.Fatal(err)
	}
	err = s.cfg.SetUser(cmd.args[1])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", cmd.args[1])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		log.Fatal("username is required")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[1],
	})
	if err != nil {
		log.Fatal(err)
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("User has been created %s\n", user.Name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		log.Fatal("Please Enter time between consecutive requests e.g. 1m: 1 min;1s,2h45m")
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
	err = scrapeFeeds(s)
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			log.Print(err)
			continue
		}
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 3 {
		log.Fatal("Please Enter name and URL of the feed")
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[1],
		Url:       cmd.args[2],
		UserID:    user.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name of the feed: %s\nName of the current User: %s\n", feedFollow.FeedName, feedFollow.UserName)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, feed := range feeds {
		feedUserName, err := s.db.GetUserFromFeed(context.Background(), feed.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name of the feed: %s\nURL of the feed: %s\nCreator of the feed: %s\n", feed.Name, feed.Url, feedUserName)
		fmt.Println()
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		log.Fatal("URL is required")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[1])
	if err != nil {
		log.Fatal(err)
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name of the feed: %s\nName of the current User: %s\n", feedFollow.FeedName, feedFollow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		log.Fatal(err)
	}
	for _, feed := range feedFollows {
		fmt.Println(feed.FeedName)
	}
	return nil
}

func handlerUnfollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		log.Fatal("Enter the URL to unfollow feed")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[1])
	if err != nil {
		log.Fatal(err)
	}
	err = s.db.DeleteUniqueFeedFollow(context.Background(), database.DeleteUniqueFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
