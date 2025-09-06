package main

import (
	"errors"
	"fmt"

	"github.com/Cheemx/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login handler expects a string argument")
	}
	err := s.cfg.SetUser(cmd.args[1])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", cmd.args[1])
	return nil
}
