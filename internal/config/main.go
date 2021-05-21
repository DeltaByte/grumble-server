package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/configor"
)

const (
	sentryDSN = "https://cf2ab432ed6349babe02271dba283cb4@o611189.ingest.sentry.io/5748077"
)

type Config struct {
	Paths  pathsConfig
	Sentry sentryConfig
	Backup backupConfig
	Port   uint   `default:"8080"`
	Host   string `default:"0.0.0.0"`
}

type pathsConfig struct {
	Database string `default:"./storage/database"`
	Media    string `default:"./storage/media"`
	Logs     string `default:"./storage/logs"`
	Backup   string `default:"./storage/backup"`
}

type sentryConfig struct {
	Enabled bool   `default:"true"`
	DSN     string `default:""`
}

type backupConfig struct {
	Schedule time.Duration `default:"6h"`
	Amount   uint16        `default:"28"`
	Group    bool          `default:"false"`
}

func Load() *Config {
	config := &Config{}

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