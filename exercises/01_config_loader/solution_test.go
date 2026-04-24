package main

import (
	"os"
	"testing"
)

func TestLoadConfig_Defaults(t *testing.T) {
	os.Clearenv()

	os.Setenv("DB_URL", "postgres://localhost")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Port != "8080" {
		t.Fatalf("expected 8080, got %s", cfg.Port)
	}

	if cfg.Debug != false {
		t.Fatalf("expected false")
	}
}

func TestLoadConfig_MissingDBURL(t *testing.T) {
	os.Clearenv()

	_, err := LoadConfig()

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestLoadConfig_DebugTrue(t *testing.T) {
	os.Clearenv()

	os.Setenv("DB_URL", "postgres://localhost")
	os.Setenv("DEBUG", "true")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Debug != true {
		t.Fatalf("expected true")
	}
}
