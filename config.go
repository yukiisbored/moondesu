package main

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type server struct {
	Port int
}

type feed struct {
	URLs           []string      `toml:"feeds"`
	UpdateDuration time.Duration `toml:"updateDuration"`
}

type config struct {
	Server server
	Feed   feed `toml:"subscription"`
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
