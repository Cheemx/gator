package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Cheemx/gator/internal/config"
	"github.com/Cheemx/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Panic(err)
	}

	db, err := sql.Open("postgres", conf.DBURL)
	if err != nil {
		log.Panic(err)
	}

	dbQueries := database.New(db)
	st := state{
		db:  dbQueries,
		cfg: &conf,
	}

	cmds := commands{
		cmdFuncMap: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	args := os.Args
	cmds.run(&st, command{name: args[1], args: args[1:]})
}
