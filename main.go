package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/FuzzyStatic/blizzard/v3"
)

func main() {
	client_id := flag.String("id", "", "Battle.net Client ID")
	client_secret := flag.String("secret", "", "Battle.net Client Secret")
	flag.Parse()

	ctx := context.Background()
	blizzClient, client_err := blizzard.NewClient(blizzard.Config{
		ClientID:     *client_id,
		ClientSecret: *client_secret,
		HTTPClient:   http.DefaultClient,
		Region:       blizzard.EU,
		Locale:       blizzard.EnGB,
	})
	if client_err != nil {
		fmt.Println(client_err)
	}

	token_err := blizzClient.AccessTokenRequest(ctx)
	if token_err != nil {
		fmt.Println(token_err)
	}

	data, _, request_err := blizzClient.WoWCharacterProfileSummary(ctx, "outland", "puzzle")
	if request_err != nil {
		fmt.Println("Error querying API", request_err)
	}

	fmt.Printf("%+v\n", data)
}
