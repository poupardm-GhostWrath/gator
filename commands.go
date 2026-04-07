package main

import (
	"errors"
)

type command struct {
	Name	string
	Args	[]string
}

type commands struct {
	list	map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.list[cmd.Name]
	if !ok {
		return errors.New("invalid command")
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(* state, command) error) {
	c.list[name] = f
}