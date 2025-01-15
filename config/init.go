// https://cobra.dev/

package config

import (
	"log"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var config *Configuration

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func initCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "hf",
		Short: "hf is habit formation",
		Long:  `hf is a habit formation offspring of the habit formation project`,
		Run: runHelp,
	}

	flags := rootCmd.Flags()
	flags.StringVar(&config.ConfigFile, "config", "./hf.json", "配置文件")

	return rootCmd
}

func ParseConfig() {
	config = &Configuration{}

	var rootCmd = initCommand()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := config.loadData(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}