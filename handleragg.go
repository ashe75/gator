package main

import (
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("time arg not found")
	}
	timeBeetweenReqsString := cmd.args[0]
	timeBeetweenReqs, err := time.ParseDuration(timeBeetweenReqsString)
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", timeBeetweenReqs)

	ticker := time.NewTicker(timeBeetweenReqs)

	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

	return nil
}
