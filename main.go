package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/Vikuuu/gator/internal/config"
	"github.com/Vikuuu/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	// Apply migrations
	if err := applyMigrations(db); err != nil {
		log.Fatalf("error applying migrations: %v", err)
	}

	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{
		registeredCommands: make(map[string]commandInfo),
	}

	cmds.register("help", func(s *state, cmd command) error {
		return handlerHelp(cmds, s, cmd)
	}, "help", "Displays all available commands and their usage")
	cmds.register("login", handlerLogin, "login <username>", "Logs a user into the system")
	cmds.register("register", handlerRegister, "register <username>", "Registers a new user")
	// cmds.register("reset", handlerReset, "reset <username>", "Resets the database")
	cmds.register("users", handlerUsers, "users", "Lists all the users in the local environment")
	cmds.register("agg", handlerAgg, "agg <time>", "Aggregates data")
	cmds.register(
		"addfeed",
		middlewareLoggedIn(handlerAddFeed),
		"addfeed <name> <url>",
		"Adds a new RSS feed",
	)
	cmds.register("feeds", handlerFeeds, "feeds", "Lists available fedds")
	cmds.register(
		"follow",
		middlewareLoggedIn(handlerFollow),
		"follow <feed_url>",
		"Follows a feed",
	)
	cmds.register(
		"following",
		middlewareLoggedIn(handlerFollowing),
		"following",
		"Lists followed feeds",
	)
	cmds.register(
		"unfollow",
		middlewareLoggedIn(handlerUnfollow),
		"unfollow <feed_url>",
		"Unfollows a feed",
	)
	cmds.register(
		"browse",
		handlerBrowse,
		"browse <number>[OPTIONAL]",
		"Browses feed content, if number given shown that number of posts else defaults to 2 posts",
	)

	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
