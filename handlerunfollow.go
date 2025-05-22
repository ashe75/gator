package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/ashe75/gator/internal/database"
	"github.com/google/uuid"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("url not found")
	}
	url := cmd.args[0]
	ctx := context.Background()
	userID := user.ID
	nullUserID := uuid.NullUUID{
		UUID:  userID,
		Valid: true,
	}

	DeleteFeedFollowByUserUrlParams := database.DeleteFeedFollowByUserUrlParams{
		UserID: nullUserID,
		Url:    url,
	}

	err := s.db.DeleteFeedFollowByUserUrl(ctx, DeleteFeedFollowByUserUrlParams)
	if err != nil {
		return err
	}
	fmt.Println("user unfollowed successfully")

	return nil
}
