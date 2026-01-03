package cmd

import (
	"bytes"
	"encoding/json"
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
	RefreshToken string
}

const (
	failedExecution int    = 1
	configName      string = ".goura"
	configExt       string = "yaml"
	apiBaseURL      string = "https://api.ouraring.com"
	userAgent       string = "goura/2.0.0"
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
		Short:         "goura is an API client for Oura Cloud API v2",
		Long:          "goura is an Unofficial API client of Oura Cloud written in Go.\nSupports Oura API v2 with all available endpoints.\nComplete documentation is available at https://github.com/paveg/goura",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cobra.OnInitialize(initConfig)

	// Existing commands
	versionCommand := versionCommand()
	configCommand := configCommand()
	userInfoCommand := userInfoCommand()
	sleepCommand := sleepCommand()
	sleepPeriodsCommand := sleepPeriodsCommand()
	sleepTimeCommand := sleepTimeCommand()
	activityCommand := activityCommand()
	readinessCommand := readinessCommand()

	// New v2 commands
	heartrateCommand := heartrateCommand()
	spo2Command := spo2Command()
	stressCommand := stressCommand()
	resilienceCommand := resilienceCommand()
	workoutCommand := workoutCommand()
	sessionCommand := sessionCommand()
	ringCommand := ringCommand()

	configCommand.Flags().StringVarP(&Config.RedirectURL, "redirectURL", "r", "http://localhost:8989", "redirect URL")

	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	// Commands that use date period flags
	dateCommands := []*cobra.Command{
		sleepCommand,
		sleepPeriodsCommand,
		sleepTimeCommand,
		activityCommand,
		readinessCommand,
		spo2Command,
		stressCommand,
		resilienceCommand,
		workoutCommand,
		sessionCommand,
		ringCommand,
	}

	for _, cmd := range dateCommands {
		cmd.Flags().StringVarP(&reqDate.target, "target", "t", "", "fetch data for a specific day (YYYY-MM-DD)")
		cmd.Flags().StringVarP(&reqDate.end, "end", "e", now.Format(dateFormat), "end date (YYYY-MM-DD)")
		cmd.Flags().StringVarP(&reqDate.start, "start", "s", lastMonth.Format(dateFormat), "start date (YYYY-MM-DD)")
	}

	// All commands to register
	allCommands := []*cobra.Command{
		versionCommand,
		configCommand,
		userInfoCommand,
		sleepCommand,
		sleepPeriodsCommand,
		sleepTimeCommand,
		activityCommand,
		readinessCommand,
		heartrateCommand,
		spo2Command,
		stressCommand,
		resilienceCommand,
		workoutCommand,
		sessionCommand,
		ringCommand,
	}

	for _, cmd := range allCommands {
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
	viper.SetDefault("ClientID", os.Getenv("OURA_CLIENT_ID"))
	viper.SetDefault("ClientSecret", os.Getenv("OURA_CLIENT_SECRET"))
	viper.SetDefault("AccessToken", os.Getenv("OURA_ACCESS_TOKEN"))
	viper.SetDefault("RedirectURL", Config.RedirectURL)
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

func out(model interface{}) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(model)
	fmt.Println(buf.String())
}
