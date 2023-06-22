package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gives the current version of the tool",
	Long:  `Current version of the tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
