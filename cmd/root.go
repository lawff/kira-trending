package cmd

import (
	"github.com/spf13/cobra"
)

var (
	kiraTrendingCmd = &cobra.Command{
		Use: "kira-trending",
	}
)

func Execute() error {
	return kiraTrendingCmd.Execute()
}
