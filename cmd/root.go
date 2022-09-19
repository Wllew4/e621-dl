package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "e621-dl",
	Long: "Download posts, pools, and queries in bulk from e621.net.",
}

func init() {
	rootCmd.AddCommand(postCmd)
	rootCmd.AddCommand(poolCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
