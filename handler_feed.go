package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Vikuuu/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.GetUserRow) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Usage: %s <name> <url>", cmd.Name)
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	// connect the feed to the user.
	userFeed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      feedName,
			Url:       feedURL,
			UserID:    user.ID,
		},
	)
	// create feed follow for the user.
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    userFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("error while adding the feed to following: %w", err)
	}

	// Print out the fields of the new feed record.
	fmt.Printf("Added Feed: %+v", userFeed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("  Name     : %s\n", feed.Name)
		fmt.Printf("  URL      : %s\n", feed.Url)
		userName, err := s.db.GetUserNameFromID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting name: %w", err)
		}
		fmt.Printf("  Username : %s\n\n", userName)
	}

	return nil
}
