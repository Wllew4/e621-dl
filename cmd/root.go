package cmd

import (
	"e621dl/dl"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "e621-dl",
	Long: "Download posts, pools, and queries in bulk from e621.net.",
}

func init() {
	rootCmd.AddCommand(postCmd)
	rootCmd.AddCommand(poolCmd)
	rootCmd.AddCommand(queryCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	start := time.Now()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	elapsed := time.Since(start)
	fmt.Printf("Downloaded %d file(s) in %s\n", dl.GetDownloadTotal(), elapsed.String())
	os.Exit(0)
}
