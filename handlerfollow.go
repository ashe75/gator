package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ashe75/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("url not found")
	}
	url := cmd.args[0]
	ctx := context.Background()

	/*currentUser := s.config.CurrentUserName
	userinf, err := s.db.GetUser(ctx, currentUser)
	if err != nil {
		return err
	} */
	userID := user.ID
	nullUserID := uuid.NullUUID{
		UUID:  userID,
		Valid: true,
	}

	id := uuid.New()
	created_at, updated_at := time.Now(), time.Now()

	feedinf, err := s.db.GetFeedsByURL(ctx, url)
	if err != nil {
		return err
	}
	feedID := feedinf.ID
	nullFeedID := uuid.NullUUID{
		UUID:  feedID,
		Valid: true,
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		UserID:    nullUserID,
		FeedID:    nullFeedID,
	}

	feedFollow, err := s.db.CreateFeedFollow(ctx, createFeedFollowParams)
	if err != nil {
		return err
	}
	fmt.Println(feedFollow)

	return nil
}
