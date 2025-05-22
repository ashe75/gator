package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ashe75/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("name and url not found")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	ctx := context.Background()

	//currentUser := s.config.CurrentUserName
	/*userinf, err := s.db.GetUser(ctx, currentUser)
	if err != nil {
		return err
	} */
	userID := user.ID
	nullUserID := uuid.NullUUID{
		UUID:  userID,
		Valid: true,
	}

	feedId := uuid.New()
	created_at, updated_at := time.Now(), time.Now()

	createFeedParams := database.CreateFeedParams{
		ID:        feedId,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		Name:      name,
		Url:       url,
		UserID:    nullUserID,
	}

	feed, err := s.db.CreateFeed(ctx, createFeedParams)
	if err != nil {
		return err
	}
	fmt.Println(feed)

	feedFollowId := uuid.New()

	nullFeedID := uuid.NullUUID{
		UUID:  feedId,
		Valid: true,
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        feedFollowId,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		UserID:    nullUserID,
		FeedID:    nullFeedID,
	}

	_, err = s.db.CreateFeedFollow(ctx, createFeedFollowParams)
	if err != nil {
		return err
	}

	return nil
}
