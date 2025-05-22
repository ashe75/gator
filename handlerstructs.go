package main

import (
	"github.com/ashe75/gator/internal/config"
	"github.com/ashe75/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}
