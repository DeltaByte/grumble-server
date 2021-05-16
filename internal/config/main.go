package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

const (
	sentryDSN = "https://cf2ab432ed6349babe02271dba283cb4@o611189.ingest.sentry.io/5748077"
)

type coreConfig struct {
	Storage storageConfig
	Sentry  sentryConfig
	Port    uint   `default:"8080"`
	Host    string `default:"0.0.0.0"`
}

type storageConfig struct {
	Database string `default:"./storage/database"`
	Media    string `default:"./storage/media"`
	Logs     string `default:"./storage/logs"`
}

type sentryConfig struct {
	Enable bool   `default:"true"`
	DSN    string `default:""`
}

func Load() *coreConfig {
	config := &coreConfig{}

	// load constants, TBH this is mostly so that longer values can be moved out
	// of the main struct so that it is slightly easier to read.
	config.Sentry.DSN = sentryDSN

	// parse env vars and load config file
	configPath := flag.String("config-file", "config.json", "Configuration file location")
	flag.Parse()

	// parse config
	err := configor.New(&configor.Config{
		ENVPrefix: "GRUMBLE",
		ErrorOnUnmatchedKeys: true,
	}).Load(config, *configPath)

	// log error and kill server if config is invalid
	if err != nil {
		fmt.Printf("Failed to load config from %s: %s", *configPath, err)
		os.Exit(1)
	}

	return config
}