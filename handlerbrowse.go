package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/ashe75/gator/internal/database"
	"github.com/google/uuid"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		val, err := strconv.Atoi(cmd.args[0])
		if err == nil {
			limit = val
		}
	}
	ctx := context.Background()

	userID := user.ID
	nullUserID := uuid.NullUUID{
		UUID:  userID,
		Valid: true,
	}
	nullLimit := sql.NullInt32{
		Int32: int32(limit),
		Valid: true,
	}
	PostUserParams := database.GetPostsForUserParams{
		UserID: nullUserID,
		Limit:  nullLimit.Int32,
	}
	posts, err := s.db.GetPostsForUser(ctx, PostUserParams)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("%v\n", post)
	}

	return nil
}
