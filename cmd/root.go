package cmd

import (
	"fmt"
	"os"
	"time"

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

type requiredDate struct {
	start  string
	end    string
	target string
}

var reqDate = &requiredDate{}

const dateFormat = "2006-01-02"

// NewCommandRoot initializes a root command function.
func NewCommandRoot() *cobra.Command {
	command := &cobra.Command{
		Use:           "goura",
		Short:         "goura is an API client of Oura Cloud",
		Long:          "goura is a Unofficial API client of Oura Cloud written in Go.\nComplete documentation is available at https://github.com/paveg/goura",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cobra.OnInitialize(initConfig)

	versionCommand := versionCommand()
	configCommand := configCommand()
	userInfoCommand := userInfoCommand()
	sleepCommand := sleepCommand()
	activityCommand := activityCommand()
	readinessCommand := readinessCommand()

	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)
	commands := []*cobra.Command{sleepCommand, activityCommand, readinessCommand}
	for _, cmd := range commands {
		cmd.Flags().StringVarP(&reqDate.target, "target", "t", "", "wanna get a specific day")
		cmd.Flags().StringVarP(&reqDate.end, "end", "e", now.Format(dateFormat), "required end date")
		cmd.Flags().StringVarP(&reqDate.start, "start", "s", lastMonth.Format(dateFormat), "required start date")
	}

	fullCommands := []*cobra.Command{versionCommand, configCommand, userInfoCommand, sleepCommand, activityCommand, readinessCommand}
	for _, cmd := range fullCommands {
		command.AddCommand(cmd)
	}

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

func initDate() (string, string, error) {
	if reqDate.target != "" {
		reqDate.start = reqDate.target
		reqDate.end = reqDate.target
	}

	startDate, err := time.Parse(dateFormat, reqDate.start)
	if err != nil {
		return "", "", err
	}
	endDate, err := time.Parse(dateFormat, reqDate.end)
	if err != nil {
		return "", "", err
	}

	return startDate.Format(dateFormat), endDate.Format(dateFormat), err
}
