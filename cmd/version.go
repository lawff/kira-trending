package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kira-trending",
	Long:  `All software has versions. This is kira-trending's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1")
	},
	Args: cobra.NoArgs,
}

func init() {
	kiraTrendingCmd.AddCommand(versionCmd)
}
