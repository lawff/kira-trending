package main

import (
	"os"

	"github.com/lawff/kira-trending/cmd"
)

// Go 程序的默认入口函数(主函数).
func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
