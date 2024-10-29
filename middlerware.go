package main

import (
	"context"
	"fmt"

	"github.com/Vikuuu/gator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *state, cmd command, user database.GetUserRow) error,
) func(*state, command) error {
	return func(s *state, cmd command) error {
		userName := s.cfg.CurrentUserName
		user, err := s.db.GetUser(context.Background(), userName)
		if err != nil {
			return fmt.Errorf("error get current user id: %w", err)
		}

		return handler(s, cmd, user)
	}
}
