package main

import (
	"fmt"
	"github.com/roderik333/graphapiauthgo"
)

func main() {
	authToken := graphapiauthgo.Auth()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Printf("Successfully authenticated: %s", authToken.Token)
}
