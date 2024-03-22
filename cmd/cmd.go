package main

import (
	"fmt"
	"github.com/roderik333/graphapiauthgo"
)

func main() {
	authToken := graphapiauthgo.Auth()
	fmt.Printf("Successfully authenticated: %s", authToken.Token)
}
