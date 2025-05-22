package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login not found or invalid")
	}
	login := cmd.args[0]
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, login)
	if err != nil {
		fmt.Println("username doesn't exists")
		os.Exit(1)
	}
	err = s.config.SetUser(login)
	if err != nil {
		return err
	}
	fmt.Printf("user %s was set succesfully\n", login)
	return nil
}
