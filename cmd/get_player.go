package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/callrua/wow-armory/utility"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getPlayerCmd)
}

var getPlayerCmd = &cobra.Command{
	Use:   "get-player",
	Short: "Print player info",
	Long:  `Print all info of a players armory`,
	Run: func(cmd *cobra.Command, args []string) {
		getPlayer()
	},
}

func getPlayer() {
	blizzClient, ctx := utility.DoAuth(&clientId, &clientSecret)
	data, _, requestErr := blizzClient.WoWCharacterProfileSummary(ctx, "outland", "puzzle")
	if requestErr != nil {
		fmt.Println("Error querying API", requestErr)
	}

	fmt.Println(prettyPrint(data))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
