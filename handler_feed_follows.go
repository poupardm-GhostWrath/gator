package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/poupardm-GhostWrath/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	// Check Args Length
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	// Assign Args to Variables
	url := cmd.Args[0]

	// Get Feed By URL
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	// Create Feed Follow
	feedFollow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:			uuid.New(),
			CreatedAt:	time.Now().UTC(),
			UpdatedAt:	time.Now().UTC(),
			UserID:		user.ID,
			FeedID:		feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Follow created successfully!")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:		%s\n", username)
	fmt.Printf("* Feed:		%s\n", feedname)
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {

	// Get Feed Follows
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf("* %s\n", feedFollow.FeedName)
	}

	return nil
}