package main

import (
	"log"
	"os"

	"github.com/Cheemx/gator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Panic(err)
	}
	st := state{
		cfg: &conf,
	}
	cmds := commands{
		cmdFuncMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		log.Println("the application expects at least 2 arguments for command and username")
		os.Exit(1)
	}
	if len(args) < 3 {
		log.Println("username is required")
		os.Exit(1)
	}
	cmds.run(&st, command{name: args[1], args: args[1:]})
}
