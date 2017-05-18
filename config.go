package main

import (
	"github.com/BurntSushi/toml"
)

type BotConfig struct {
	Token string
}

func loadConfig() (out BotConfig, err error) {
	_, err = toml.DecodeFile("config.toml", &out)
	return
}
