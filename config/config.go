package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type (
	ConfigList struct {
		ApiKey    string
		ApiSecret string
		TradeSize float64
		SlackURL  string
	}
)

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file:config.ini, err: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
		TradeSize: cfg.Section("trade").Key("size").MustFloat64(),
		SlackURL:  cfg.Section("slack").Key("url").String(),
	}
}
