package main

import "log"

func main() {
	log.Println("starting moondesu")

	_, err := loadConfiguration("./config.toml")

	if err != nil {
		log.Panic(err)
	}
}
