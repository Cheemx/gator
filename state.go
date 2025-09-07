package main

import (
	"context"
	"fmt"
	"log"
	"os"
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
		log.Println("username is required")
		os.Exit(1)
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
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *feed)
	return nil
}
