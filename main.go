package main

import (
	"log"
	"os"
	"github.com/poupardm-GhostWrath/gator/internal/config"
)

type state struct {
	cfg	*config.Config
}

func main() {
	// Read Config File
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Store cfg in state struct
	programState := &state{
		cfg: &cfg,
	}

	// Create commands struct and register login handler
	cmds := commands{
		list: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

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