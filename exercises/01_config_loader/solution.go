package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port  string
	DBURL string
	Debug bool
}

func LoadConfig() (Config, error) {
	cfg := Config{}

	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}

	cfg.Port = port

	dbURL := strings.TrimSpace(os.Getenv("DB_URL"))
	if dbURL == "" {
		return Config{}, errors.New("DB_URL required")
	}
	cfg.DBURL = dbURL

	debugRaw := strings.TrimSpace(os.Getenv("DEBUG"))
	if debugRaw != "" {
		val, err := strconv.ParseBool(debugRaw)
		if err != nil {
			return Config{}, errors.New("invalid debug value")
		}
		cfg.Debug = val
	}
	return cfg, nil
}
