package main

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

type BaseSettings struct {
	ClientSecret string
	ClientID     string
	TenantID     string
	ApiEndpoint  string
	Scope        []string
}

type Settings struct {
	BaseDir string
	Msal    BaseSettings
}

func NewBaseSettings() BaseSettings {
	return BaseSettings{
		ClientSecret: getEnv("client_secret", ""),
		ClientID:     getEnv("client_id", ""),
		TenantID:     getEnv("tenant_id", ""),
		ApiEndpoint:  getEnv("api_endpoint", "https://graph.microsoft.com/beta"),
		Scope:        []string{getEnv("scope", "https://graph.microsoft.com/.default")},
	}
}

func NewSettings() Settings {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return Settings{
		BaseDir: dir,
		Msal:    NewBaseSettings(),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
