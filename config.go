package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type server struct {
	Port int
}

type feed struct {
	Name    string
	Website string
	Feed    string
}

type config struct {
	Server server
	Feeds  []feed `toml:"feed"`
}

func loadConfiguration(path string) (*config, error) {
	cfg := new(config)

	_, err := toml.DecodeFile(path, cfg)
	if err != nil {
		err = fmt.Errorf("error parsing configuration file: %v", err)
		return nil, err
	}

	return cfg, nil
}
