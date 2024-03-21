package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Settings struct {
	ClientID     string
	ClientSecret string
	TenantID     string
	Scope        string
}

func FetchAuthToken(settings Settings) (string, error) {
	config := &clientcredentials.Config{
		ClientID:     settings.ClientID,
		ClientSecret: settings.ClientSecret,
		TokenURL:     fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", settings.TenantID),
		Scopes:       []string{settings.Scope},
	}

	ctx := context.Background()
	token, err := config.Token(ctx)
	if err != nil {
		return "", fmt.Errorf("authentication failed: %w", err)
	}

	return token.AccessToken, nil
}

func main() {
	godotenv.Load()
	Settings := NewSettings()
	// Use settings
	settings := Settings{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		TenantID:     "your-tenant-id",
		Scope:        "your-scope",
	}

	token, err := FetchAuthToken(settings)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Access token:", token)
}
