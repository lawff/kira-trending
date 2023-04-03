package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

var (
	keyword string
)

var githubSCmd = &cobra.Command{
	Use:   "gs",
	Short: "Search GitHub repositories",
	RunE:  searchRepo,
	PostRun: func(cmd *cobra.Command, args []string) {
		// print the limit
		fmt.Println("\nThat's all.")
	},
	SuggestFor: []string{"ga"},
	Args:       cobra.ExactArgs(0),
}

func searchRepo(cmd *cobra.Command, args []string) error {
	if keyword == "" {
		return fmt.Errorf("missing required flag: keyword")
	}

	fmt.Println("\nSearching...")
	fmt.Println()

	ctx := context.Background()
	client := github.NewClient(nil)

	query := fmt.Sprintf("%s in:name,description", keyword)
	opt := &github.SearchOptions{
		Sort:        "stars",
		Order:       "desc",
		ListOptions: github.ListOptions{PerPage: 20},
	}

	result, _, err := client.Search.Repositories(ctx, query, opt)
	if err != nil {
		return err
	}

	fmt.Printf("%d repository results:\n", result.GetTotal())
	for i, repo := range result.Repositories {
		// 还需要有链接
		fmt.Printf("%d. %s (🌟 %d) [%s]\n", i+1, repo.GetFullName(), repo.GetStargazersCount(), repo.GetHTMLURL())
	}

	return nil
}

func init() {
	githubSCmd.PersistentFlags().StringVarP(&keyword, "keyword", "k", "", "search keyword (required)")
	kiraTrendingCmd.AddCommand(githubSCmd)
}
