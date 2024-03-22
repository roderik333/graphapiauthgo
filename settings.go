package graphapiauthgo

import (
	"os"
	"path/filepath"
)

type baseSettings struct {
	ClientSecret string
	ClientID     string
	TenantID     string
	ApiEndpoint  string
	Scope        []string
}

type Settings struct {
	BaseDir string
	Msal    baseSettings
}

func getBaseSettings() baseSettings {
	return baseSettings{
		ClientSecret: getEnv("client_secret", ""),
		ClientID:     getEnv("client_id", ""),
		TenantID:     getEnv("tenant_id", ""),
		ApiEndpoint:  getEnv("api_endpoint", "https://graph.microsoft.com/beta"),
		Scope:        []string{getEnv("scope", "https://graph.microsoft.com/.default")},
	}
}

func GetSettings() Settings {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return Settings{
		BaseDir: dir,
		Msal:    getBaseSettings(),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
