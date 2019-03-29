// Copyright © 2019 Lawrence Livermore National Security
// See LICENSE file

package cmd

import (
	"fmt"
	"log"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var confFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command {
	Use: "gotls",
	Short: "gotls is an automated TLS certificate issuance and management tool",
	Long: `gotls can generate keys, CSRs, and optionally obtain the certificate
with an internal Active Directory Certificate Services endpoint or the
Let's Encrypt service.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// persistent flags will be global for your application.
	rootCmd.PersistentFlags().StringVar(&confFile, "config", "", "config file (default is $HOME/.gotls.yaml)")

	// local flags will only run when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if confFile != "" {
		viper.SetConfigFile(confFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		viper.AddConfigPath(home)
		viper.AddConfigPath(".") // override config from the working directory
		viper.SetConfigName(".gotls") // do not set the extension; viper will try all known types
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("using config file: %s\n\n", viper.ConfigFileUsed())
	}
}
