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
	fmt.Println("Feed created successfully!")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("=====================================")

	err = handlerFollow(
		s, 
		command{
			Name: "follow",
			Args: []string{feed.Url},
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}
	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID: 		%v\n", feed.ID)
	fmt.Printf("* Created:	%v\n", feed.CreatedAt)
	fmt.Printf("* Updated:	%v\n", feed.UpdatedAt)
	fmt.Printf("* Name:		%v\n", feed.Name)
	fmt.Printf("* URL:		%v\n", feed.Url)
	fmt.Printf("* User:		%v\n", user.Name)
}

func handlerListFeeds(s *state, cmd command) error {
	// Get Feeds from DB
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		printFeed(feed, user)
		fmt.Println("=====================================")
	}
	return nil
}