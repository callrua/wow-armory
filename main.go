package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/FuzzyStatic/blizzard/v3"
	"github.com/callrua/wow-armory/utility"
	"net/http"
)

func main() {
	utility.GetOauth("hi", "you")
	clientId := flag.String("id", "", "Battle.net Client ID")
	clientSecret := flag.String("secret", "", "Battle.net Client Secret")
	flag.Parse()

	ctx := context.Background()
	blizzClient, clientErr := blizzard.NewClient(blizzard.Config{
		ClientID:     *clientId,
		ClientSecret: *clientSecret,
		HTTPClient:   http.DefaultClient,
		Region:       blizzard.EU,
		Locale:       blizzard.EnGB,
	})
	if clientErr != nil {
		fmt.Println(clientErr)
	}

	tokenErr := blizzClient.AccessTokenRequest(ctx)
	if tokenErr != nil {
		fmt.Println(tokenErr)
	}

	data, _, requestErr := blizzClient.WoWCharacterProfileSummary(ctx, "outland", "puzzle")
	if requestErr != nil {
		fmt.Println("Error querying API", requestErr)
	}

	fmt.Printf("%+v\n", data)
}
