package main

import (
	"fmt"
	"time"
	"context"

	"github.com/google/uuid"
	"github.com/poupardm-GhostWrath/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	// Check Args Length
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	
	// Assign Args to Variables
	name := cmd.Args[0]
	url := cmd.Args[1]

	// Get User From DB
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	// Create Feed
	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:       	uuid.New(),
			CreatedAt:	time.Now().UTC(),
			UpdatedAt: 	time.Now().UTC(),
			Name:      	name,
			Url:       	url,
			UserID:    	user.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	fmt.Println("New feed created successfully!")
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID: 		%v\n", feed.ID)
	fmt.Printf(" * Name: 	%v\n", feed.Name)
	fmt.Printf(" * URL: 	%v\n", feed.Url)
	fmt.Printf(" * User ID:	%v\n", feed.UserID)
}