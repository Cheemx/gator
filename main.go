package main

import (
	"database/sql"
	"embed"
	"log"
	"os"

	"github.com/Cheemx/gator/internal/config"
	"github.com/Cheemx/gator/internal/database"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed sql/schema/*.sql
var embedMigrations embed.FS

func initDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	if err := goose.Up(db, "sql/schema"); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Panic(err)
	}

	db, err := initDB(conf.DBURL)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

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
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollowing))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	args := os.Args
	cmds.run(&st, command{name: args[1], args: args[1:]})
}
