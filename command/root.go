package command

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

const (
	failedExecution int = 1
)

// NewCommandRoot initializes a root command function.
func NewCommandRoot() *cobra.Command {
	command := &cobra.Command{
		Use:   "goura",
		Short: "goura is a API client of Oura Cloud",
		Long:  "goura is a Unofficial API client of Oura Cloud written in Go.\nComplete documentation is available at https://github.com/paveg/goura",
	}
	cobra.OnInitialize(initConfig)

	command.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.goura.yaml)")

	command.AddCommand(versionCommand())
	return command
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(failedExecution)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".goura")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %#v", viper.ConfigFileUsed())
	}
}

// Execute executes NewCommandRoot function.
func Execute() {
	command := NewCommandRoot()
	command.SetOut(os.Stdout)
	if err := command.Execute(); err != nil {
		command.SetOut(os.Stderr)
		command.Println(err)
		os.Exit(failedExecution)
	}
}
