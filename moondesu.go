package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var configPath string

func processArguments() {
	flag.StringVar(&configPath,
		"conf", "./config.toml", "Location of the config file")

	flag.Parse()
}

func addFeedsToSubscription(cfg *config) error {
	for _, url := range cfg.Feed.URLs {
		log.Printf("adding %s to subscription\n", url)

		err := updateSubscription(url)
		if err != nil {
			return fmt.Errorf("error while subscribing %s", err)
		}
	}

	return nil
}

func main() {
	processArguments()

	log.Println("starting moondesu")

	config, err := loadConfiguration(configPath)
	if err != nil {
		log.Panic(err)
	}

	log.Println("adding feeds to subscription")
	err = addFeedsToSubscription(config)
	if err != nil {
		log.Panic(err)
	}

	log.Println("starting subscription ticker")
	startSubcriptionTicker(config.Feed.UpdateDuration * time.Second)
}
