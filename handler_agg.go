package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Vikuuu/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching link: %w", err)
	}

	fmt.Printf("RSS Feed:  %+v", rssFeed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Usage: %s <name> <url>", cmd.Name)
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	// Get the current User name.
	userName := s.cfg.CurrentUserName
	// Get the current User ID.
	userId, err := s.db.GetUserID(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error get current user id: %w", err)
	}
	// connect the feed to the user.
	userFeed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      feedName,
			Url:       feedURL,
			UserID:    userId,
		},
	)

	// Print out the fields of the new feed record.
	fmt.Printf("Added Feed: %+v", userFeed)
	return nil
}
