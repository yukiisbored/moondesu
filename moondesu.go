package main

import (
	"flag"
	"log"
)

var configPath string

func processArguments() {
	flag.StringVar(&configPath,
		"conf", "./config.toml", "Location of the config file")

	flag.Parse()
}

func main() {
	processArguments()

	log.Println("starting moondesu")

	_, err := loadConfiguration(configPath)

	if err != nil {
		log.Panic(err)
	}
}
