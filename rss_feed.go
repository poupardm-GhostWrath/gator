package main

import (
	"net/http"
	"io"
	"encoding/xml"
	"html"
	"context"
	"time"
)

type RSSItem struct {
	Title		string	`xml:"title"`
	Link		string	`xml:"link"`
	Description	string	`xml:"description"`
	PubDate		string	`xml:"pubDate"`
}

type RSSFeed struct {
	Channel struct {
		Title		string		`xml:"title"`
		Link		string		`xml:"link"`
		Description	string		`xml:"description"`
		Item		[]RSSItem	`xml:"item"`
	} `xml:"channel"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Create Client
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	// Make New Request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	// Set User-Agent Header
	req.Header.Set("User-Agent", "gator")

	// Send Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read Data
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal Data
	var feedResp RSSFeed
	err = xml.Unmarshal(dat, &feedResp)
	if err != nil {
		return nil, err
	}

	// Decode Escaped HTML
	err = decodeHTML(&feedResp)
	if err != nil {
		return nil, err
	}

	return &feedResp, nil
}

func decodeHTML(feed *RSSFeed) error {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := 0; i < len(feed.Channel.Item); i++ {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}
	return nil
}