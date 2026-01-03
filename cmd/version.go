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
	majorVersion int = 2
	minorVersion int = 0
	patchVersion int = 0
)

func versionCommand() *cobra.Command {
	version = fmt.Sprintf("v%v.%v.%v", majorVersion, minorVersion, patchVersion)
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of goura",
		Long:  "Print the version number of goura (Oura API v2 client)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("goura version: %s, revision: %s\n", version, revision)
			fmt.Println("Oura API: v2")
		},
	}

	return cmd
}

func init() {
}
