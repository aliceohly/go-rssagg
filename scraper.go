package main

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/aliceohly/go-rssagg/internal/database"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Items       []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines", timeBetweenRequest, concurrency)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetLastFeeds(context.Background(), int32(concurrency)) // a bit confused on the context.Background() part
		if err != nil {
			log.Printf("Error getting feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed fetched: %v", err)
		return
	}

	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v", err)
		return
	}

	for _, item := range feedData.Channel.Items {
		log.Printf("Inserting item %s", item.Title)
		// TODO: insert item into database
	}

	log.Printf("Feed %s fetched, %v posts found", feed.Name, len(feedData.Channel.Items))
}

func fetchFeed(feedUrl string) (*RssFeed, error) {
	response, err := http.Get(feedUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close() // why do we need to close the body?

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RssFeed
	err = xml.Unmarshal(data, &rssFeed) // how does unmarchal work?
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}
