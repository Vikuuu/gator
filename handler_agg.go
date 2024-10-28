package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Vikuuu/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <time>", cmd.Name)
	}
	arg := cmd.Args[0]

	timeBetweenReq, err := time.ParseDuration(arg)
	if err != nil {
		return fmt.Errorf("Input time as '1h', '1m', '1s'")
	}

	fmt.Printf("Collecting Feed every %s\n", timeBetweenReq)
	ticker := time.NewTicker(timeBetweenReq)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feed ID: %w", err)
	}

	currUser := s.cfg.CurrentUserName
	currUserID, err := s.db.GetUser(context.Background(), currUser)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:     time.Now().UTC(),
		ID:            feed.ID,
		UserID:        currUserID.ID,
	})
	if err != nil {
		return fmt.Errorf("error marking feed fetched: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching the feed: %w", err)
	}

	printRssFeed(rssFeed)
	return nil
}

func printRssFeed(rssFeed *RSSFeed) {
	fmt.Printf("Title      : %s\n", rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf(" - Title 	   : %s\n", item.Title)
	}
}
