package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	feedResp := RSSFeed{}
	err = xml.Unmarshal(dat, &feedResp)
	if err != nil {
		return &RSSFeed{}, err
	}

	feedResp.Channel.Description = html.UnescapeString(feedResp.Channel.Description)
	feedResp.Channel.Title = html.UnescapeString(feedResp.Channel.Title)
	for i := range feedResp.Channel.Item {
		feedResp.Channel.Item[i].Description = html.UnescapeString(feedResp.Channel.Item[i].Description)
		feedResp.Channel.Item[i].Title = html.UnescapeString(feedResp.Channel.Item[i].Title)
	}
	return &feedResp, nil
}
