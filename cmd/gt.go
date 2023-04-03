package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
}

var limit int

var githubTCmd = &cobra.Command{
	Use:   "gt",
	Short: "View GitHub trending repositories",
	Long:  `A simple command-line tool to view GitHub trending repositories`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// check the args

	},
	Run: func(cmd *cobra.Command, args []string) {
		viewTrending(args[0])
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// print the limit
		fmt.Println("\nThat's all.")
	},
	Args:       cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	SuggestFor: []string{"gn"},
	ValidArgs: []string{
		"javascript",
		"ruby",
		"python",
		"java",
		"php",
		"css",
		"typescript",
		"c++",
		"c",
		"shell",
		"swift",
		"objective-c",
		"go",
		"scala",
		"rust",
		"coffeescript",
		"elixir",
		"clojure",
		"lua",
		"perl",
		"r",
	},
}

func viewTrending(language string) {
	fmt.Println("\nFetching...")
	url := fmt.Sprintf("https://github.com/trending/%s", language)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nTop %d trending %s repositories on GitHub:\n", limit, language)
	fmt.Println()
	doc.Find(".Box .Box-row").Each(func(i int, s *goquery.Selection) {
		today := strings.TrimSpace(s.Find(".f6 .float-sm-right").Text())
		description := strings.TrimSpace(s.Find("p").Text())
		stars := strings.TrimSpace(s.Find(".f6 .Link--muted").Eq(0).Text())
		forks := strings.TrimSpace(s.Find(".f6 .Link--muted").Eq(1).Text())
		url, exists := s.Find(".h3 a").Attr("href")
		if exists && i < limit {
			fmt.Printf("%d. %s (%s/%s) - %s\n   URL: https://github.com%s\n", i+1, url[1:], today, language, description, url)
			fmt.Printf("   Stars: %s | Forks: %s\n", stars, forks)
		}
	})

}

func init() {
	githubTCmd.Flags().IntVarP(&limit, "limit", "l", 10, "view the top N trending repositories (default 10)")
	kiraTrendingCmd.AddCommand(githubTCmd)
}
