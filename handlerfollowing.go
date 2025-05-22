package main

import (
	"context"
	"fmt"

	"github.com/ashe75/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
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
	feeds, err := s.db.GetFeedFollowsForUser(ctx, nullUserID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%v\n", feed.FeedName)
	}
	return nil
}
