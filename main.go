package main

import (
	"database/sql"
	"log"
	"os"

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