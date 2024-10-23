package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Vikuuu/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.GetUserRow) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feedId, err := s.db.GetFeedIDFromURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed name: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedId,
	})

	fmt.Printf("Feed Name: %s\n", feedFollow.FeedName)
	fmt.Printf("User Name: %s\n", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.GetUserRow) error {
	feedFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting following feeds: %w", err)
	}

	for _, feed := range feedFollowing {
		fmt.Printf("%s\n", feed.FeedName)
	}

	return nil
}
