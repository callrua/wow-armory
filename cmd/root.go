package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	// CLI flags
	clientId		string
	clientSecret	string

	rootCmd = &cobra.Command{
		Use:	"wow-armory",
		Short:	"Dump info from the World of Warcraft Armory",
		Long: 	`Multiple options to dump different things from the World of Warcraft Armory API`,
		Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&clientId, "id", "", "client id")
	rootCmd.PersistentFlags().StringVar(&clientSecret, "secret", "", "client secret")
	err := viper.BindPFlag("clientId", rootCmd.PersistentFlags().Lookup("clientId"))
	if err != nil {
		return
	}
	err = viper.BindPFlag("clientSecret", rootCmd.PersistentFlags().Lookup("clientSecret"))
	if err != nil {
		return
	}

	rootCmd.AddCommand(getPlayerCmd)
}

func initConfig() {
	viper.AutomaticEnv()
}