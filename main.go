package main

import "github.com/callrua/wow-armory/cmd"

func main() {
	//clientId := flag.String("id", "", "Battle.net Client ID")
	//clientSecret := flag.String("secret", "", "Battle.net Client Secret")
	//flag.Parse()
	//
	//blizzClient, ctx := utility.DoAuth(clientId, clientSecret)
	//data, _, requestErr := blizzClient.WoWCharacterProfileSummary(ctx, "outland", "puzzle")
	//if requestErr != nil {
	//	fmt.Println("Error querying API", requestErr)
	//}
	//
	//fmt.Printf("%+v\n", data)
	cmd.Execute()
}
