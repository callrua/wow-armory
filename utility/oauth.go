package utility

import (
	"context"
	"fmt"
	"net/http"

	"github.com/FuzzyStatic/blizzard/v3"
)

// DoAuth will make an authorization request to the Battle.net system,
// expecting a Battle.net Client ID and Client Secret as input.
func DoAuth(clientId string, clientSecret string) (*blizzard.Client, context.Context){
	ctx := context.Background()
	blizzClient, clientErr := blizzard.NewClient(blizzard.Config{
		ClientID: clientId,
		ClientSecret: clientSecret,
		HTTPClient: http.DefaultClient,
		Region:	blizzard.EU,
		Locale: blizzard.EnGB,
	})
	if clientErr != nil {
		fmt.Println(clientErr)
	}

	tokenErr := blizzClient.AccessTokenRequest(ctx)
	if tokenErr != nil {
		fmt.Println(tokenErr)
	}

	return blizzClient, ctx
}
