package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetUserNameCreatedFeed(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%v %v %v\n", feed.Name, feed.Url, feed.Name_2)
	}
	return nil
}
