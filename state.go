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
	fmt.Printf("User has been created %+v\n", user)
	return nil
}
