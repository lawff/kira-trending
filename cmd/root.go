package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	kiraTrendingCmd = &cobra.Command{
		Use: "kira-trending",
	}
)

func Execute() error {
	err := doc.GenMarkdownTree(kiraTrendingCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
	return kiraTrendingCmd.Execute()
}
