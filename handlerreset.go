package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	err := s.db.DeleteUsers(ctx)
	if err != nil {
		fmt.Println("reset was unsuccessful")
		os.Exit(1)
	}
	fmt.Println("table was reset")
	return nil
}
