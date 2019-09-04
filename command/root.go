package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	failedExecution int = 1
)

var rootCommand = &cobra.Command{
	Use:   "goura",
	Short: "goura is a API client of Oura Cloud",
	Long:  "goura is a Unofficial API client of Oura Cloud written in Go.\nComplete documentation is available at https://github.com/paveg/goura",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(failedExecution)
	}
}
