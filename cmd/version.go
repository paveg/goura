package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version  string
	revision string
)

const (
	majorVersion int = 1
	minorVersion int = 0
	patchVersion int = 0
)

func versionCommand() *cobra.Command {
	version = fmt.Sprintf("v%v.%v.%v", majorVersion, minorVersion, patchVersion)
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of goura",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("goura version: %s, revision: %s\n", version, revision)
		},
	}

	return cmd
}

func init() {
}
