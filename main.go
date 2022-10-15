package main

import (
	"fmt"

	"github.com/callrua/wow-armory/utility"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// Get Client details, prefer env var over cli flag
	var clientId, clientSecret string

	viper.SetEnvPrefix("CLIENT")
	viper.BindEnv("id")
	viper.BindEnv("secret")

	pflag.String("id", "", "Battle.net client id")
	pflag.String("secret", "", "Battle.net client secret")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	clientId = viper.GetString("id")
	clientSecret = viper.GetString("secret")

	blizzClient, ctx := utility.DoAuth(clientId, clientSecret)
	data, _, requestErr := blizzClient.WoWCharacterProfileSummary(ctx, "outland", "puzzle")
	if requestErr != nil {
		fmt.Println("Error querying API", requestErr)
	}

	fmt.Printf("%+v\n", data)
}
