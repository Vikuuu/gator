package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Vikuuu/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feedId, err := s.db.GetFeedIDFromURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed name: %w", err)
	}

	userName := s.cfg.CurrentUserName
	userId, err := s.db.GetUserID(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userId,
		FeedID:    feedId,
	})

	fmt.Printf("Feed Name: %s", feedFollow.FeedName)
	fmt.Printf("User Name: %s", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	userName := s.cfg.CurrentUserName
	userId, err := s.db.GetUserID(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error getting user id: %w", err)
	}

	feedFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), userId)
	if err != nil {
		return fmt.Errorf("error getting following feeds: %w", err)
	}

	for _, feed := range feedFollowing {
		fmt.Printf("%s", feed.FeedName)
	}

	return nil
}
