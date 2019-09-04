package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	majorVersion int = 1
	minorVersion int = 0
	patchVersion int = 0
)

func versionCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of goura",
		Long:  `All software has versions. This is goura's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("goura API client of oura cloud v%v.%v.%v", majorVersion, minorVersion, patchVersion)
		},
	}

	return command
}

func init() {
}
