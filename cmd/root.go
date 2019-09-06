package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config declares config
var Config config

// Config represents config struct
type config struct {
	Debug        bool
	RedirectURL  string
	ClientID     string
	ClientSecret string
	AccessToken  string
}

const (
	failedExecution int    = 1
	configName      string = ".goura"
	configExt       string = "yaml"
	apiBaseURL      string = "https://api.ouraring.com"
)

// NewCommandRoot initializes a root command function.
func NewCommandRoot() *cobra.Command {
	command := &cobra.Command{
		Use:           "goura",
		Short:         "goura is a API client of Oura Cloud",
		Long:          "goura is a Unofficial API client of Oura Cloud written in Go.\nComplete documentation is available at https://github.com/paveg/goura",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cobra.OnInitialize(initConfig)

	versionCommand := versionCommand()
	configCommand := configCommand()
	userInfoCommand := userInfoCommand()
	sleepCommand := sleepCommand()

	command.AddCommand(versionCommand)
	command.AddCommand(configCommand)
	command.AddCommand(userInfoCommand)
	command.AddCommand(sleepCommand)
	return command
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(failedExecution)
	}

	viper.SetConfigName(configName)
	viper.SetConfigType(configExt)
	viper.AddConfigPath(home)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		filePath := fmt.Sprintf("%s/%s.%s", home, configName, configExt)
		createConfigFile(filePath)
	}
	viper.SetDefault("RedirectURL", "http://localhost:8989")
	viper.SetDefault("ClientID", os.Getenv("OURA_CLIENT_ID"))
	viper.SetDefault("ClientSecret", os.Getenv("OURA_CLIENT_SECRET"))

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(failedExecution)
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

func createConfigFile(filePath string) {
	fmt.Printf("create config file... (%s)\n", filePath)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(failedExecution)
	}
	defer file.Close()
}
