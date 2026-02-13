package cmd

import (
	"fmt"
	"os"
	"samurenkoroma/services/configs"

	"github.com/spf13/cobra"
)

var conf *configs.Config

func init() {
	cobra.OnInitialize(initConfig)

}

var rootCmd = &cobra.Command{
	Use:   "[your-cli-app-name]",
	Short: "A brief description of your CLI application",
	Long: `A longer description that explains your CLI application in detail, 
    including available commands and their usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to [your-cli-app-name]! Use --help for usage.")
	},
}

func initConfig() {
	conf = configs.LoadConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
