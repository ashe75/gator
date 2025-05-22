package main

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/ashe75/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	lastFetchedNotNull := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	markFeedFetchParams := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: lastFetchedNotNull,
		ID:            nextFeed.ID,
	}
	feed, err := s.db.MarkFeedFetched(context.Background(), markFeedFetchParams)
	if err != nil {
		return err
	}
	RSSfeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, item := range RSSfeed.Channel.Item {
		descriptionNotNull := sql.NullString{
			String: item.Description,
			Valid:  true,
		}
		pubTime, err := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", item.PubDate)
		if err != nil {
			return err
		}
		nullFeedID := uuid.NullUUID{
			UUID:  feed.ID,
			Valid: true,
		}
		postsParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: descriptionNotNull,
			PublishedAt: pubTime,
			FeedID:      nullFeedID,
		}
		_, err = s.db.CreatePost(context.Background(), postsParams)
		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			return err
		}
	}
	return nil
}
