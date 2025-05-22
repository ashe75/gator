package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ashe75/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login not found or invalid")
	}
	name := cmd.args[0]

	ctx := context.Background()
	id := uuid.New()
	created_at, updated_at := time.Now(), time.Now()
	_, err := s.db.GetUser(ctx, name)
	if err == nil {
		fmt.Println("user already exist")
		os.Exit(1)
	}

	createUserParams := database.CreateUserParams{
		ID:        id,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		Name:      name,
	}

	user, err := s.db.CreateUser(ctx, createUserParams)
	if err != nil {
		return err
	}
	err = s.config.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("user %s was created", user.Name)
	fmt.Println(user)

	return nil
}
