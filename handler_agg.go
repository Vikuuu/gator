package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/Vikuuu/gator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	args := 2
	if len(cmd.Args) == 1 {
		l, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("error <limit> should be a number: %w", err)
		}

		args = l
	}

	posts, err := s.db.GetPostsForUser(context.Background(), args)
	if err != nil {
		return fmt.Errorf("error while getting posts for user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Pub at: %s\n", post.PublishedAt)
		fmt.Printf("Description: %s\n\n", post.Description)
	}

	return nil
}

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
		fmt.Println("Feed Scraped")
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

	saveInPosts(s, rssFeed, feed.ID)
	return nil
}

func saveInPosts(s *state, rssFeed *RSSFeed, feedID uuid.UUID) error {
	for _, item := range rssFeed.Channel.Item {
		layout := "Mon, 02 Jan 2006 15:04:05 -0700"
		t, err := time.Parse(layout, item.PubDate)
		if err != nil {
			log.Fatalf("Error Parsing time %v: ", err)
		}

		err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         sql.NullString{String: item.Link, Valid: true},
			Description: item.Description,
			PublishedAt: t,
			FeedID:      feedID,
		})
		if err != nil {
			return fmt.Errorf("error saving posts: %w", err)
		}
	}
	return nil
}
