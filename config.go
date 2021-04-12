package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

type Config struct {
	Port uint `default:"8080"`
	Host string `default:"0.0.0.0"`
}


func LoadConfig() *Config {
	config := &Config{}

	// parse env vars and load config file
	configPath := flag.String("config-file", "config.json", "Configuration file location")
	flag.Parse()

	// parse config
	err := configor.New(&configor.Config{
		ENVPrefix: "GRUMBLE",
		ErrorOnUnmatchedKeys: true,
	}).Load(config, *configPath)

	// log error and kill server in config is invalid
	if err != nil {
		fmt.Printf("Failed to load config from %s: %s", *configPath, err)
		os.Exit(1)
	}

	return config
}