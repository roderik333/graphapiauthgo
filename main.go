package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/clientcredentials"
)

func FetchAuthToken(settings Settings) (string, error) {
	config := &clientcredentials.Config{
		ClientID:     settings.Msal.ClientID,
		ClientSecret: settings.Msal.ClientSecret,
		TokenURL:     fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", settings.Msal.TenantID),
		Scopes:       settings.Msal.Scope,
	}

	ctx := context.Background()
	token, err := config.Token(ctx)
	if err != nil {
		return "", fmt.Errorf("authentication failed: %w", err)
	}

	return token.AccessToken, nil
}

type AccessToken struct {
	Token string
}

func Auth() AccessToken {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	settings := GetSettings()
	token, err := FetchAuthToken(settings)
	if err != nil {
		panic(err)
	}
	return AccessToken{Token: token}
}
