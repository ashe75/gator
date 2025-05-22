package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ashe75/gator/internal/config"
	"github.com/ashe75/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", conf.DbUrl)

	dbQueries := database.New(db)

	s := state{
		db:     dbQueries,
		config: &conf,
	}

	handlers := make(map[string]func(*state, command) error)
	c := commands{
		handlers: handlers,
	}

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", middlewareLoggedIn(handlerFollowing))
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	c.register("browse", middlewareLoggedIn(handlerBrowse))
	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough args")
		os.Exit(1)
	}
	commandName := args[1]
	commandArgs := args[2:]

	newCommand := command{
		name: commandName,
		args: commandArgs,
	}

	err = c.run(&s, newCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
