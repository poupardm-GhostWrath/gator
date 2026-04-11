package main

import (
	"database/sql"
	"log"
	"os"
	"context"

	"github.com/poupardm-GhostWrath/gator/internal/config"
	"github.com/poupardm-GhostWrath/gator/internal/database"
	_ "github.com/lib/pq"
)



type state struct {
	db	*database.Queries
	cfg	*config.Config
}

func main() {
	// Read Config File
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Open database
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	// Create Queries
	dbQueries := database.New(db)

	// Store cfg in state struct
	programState := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	// Create commands struct and register login handler
	cmds := commands{
		list: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	// Get the args and split
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
	
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}