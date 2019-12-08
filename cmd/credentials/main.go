package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zmb3/spotify"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	redirectURI := os.Getenv("SPOTIFY_REDIRECT_URI")
	auth := spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistModifyPrivate)
	state := "jesusgoku"
	url := auth.AuthURL(state)

	fmt.Printf("Open then next url on your browser and authorize app: %v\n\n", url)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter authorization code: ")
	authToken, _ := reader.ReadString('\n')
	authToken = strings.TrimSuffix(authToken, "\n")

	token, err := auth.Exchange(authToken)

	if err != nil {
		fmt.Printf("\nCannot fetch token\n\n")
		os.Exit(1)
		return
	}

	fmt.Printf("\nAdd both tokens to .env file:\n\n")
	fmt.Printf("Access Token: %v\n\n", token.AccessToken)
	fmt.Printf("Refresh Token: %v\n\n", token.RefreshToken)
}
