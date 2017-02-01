package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

var subscription = make(map[string]*gofeed.Feed)
var subscriptionTicker *time.Ticker

func fetchFeed(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching feed: %v", err)
	}

	return feed, nil
}

func fetchFeeds(urls []string) ([]*gofeed.Feed, error) {
	var feeds []*gofeed.Feed

	for _, url := range urls {
		feed, err := fetchFeed(url)
		if err != nil {
			return feeds, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}

func updateSubscription(url string) error {
	feed, err := fetchFeed(url)
	if err != nil {
		return err
	}

	subscription[url] = feed

	return nil
}

func updateSubscriptions() error {
	for k := range subscription {
		err := updateSubscription(k)
		return err
	}

	return nil
}

func startSubcriptionTicker(d time.Duration) {
	subscriptionTicker = time.NewTicker(d)

	go func() {
		for _ = range subscriptionTicker.C {
			log.Println("updating subscriptions")
			err := updateSubscriptions()
			if err != nil {
				log.Fatalf("error while updating: %v", err)
			}
			log.Println("subscriptions has been updated")
		}
	}()
}
