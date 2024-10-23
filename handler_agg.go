package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching link: %w", err)
	}

	fmt.Printf("RSS Feed:  %+v", rssFeed)
	return nil
}
