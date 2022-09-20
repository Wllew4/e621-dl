package cmd

import (
	"e621-dl/dl"
	"e621-dl/e6api"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	query_tags  []string
	query_first int
	query_out   string
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Download the first n posts from a search",
	Long:  "Download the first n posts from a search",
	Run: func(cmd *cobra.Command, args []string) {
		// fetch posts
		fmt.Println("Fetching posts")
		fmt.Println(query_tags)
		posts := e6api.GetPostsByTags(query_tags, query_first)

		os.MkdirAll(query_out, fs.ModePerm)
		for i, p := range posts {
			fmt.Println(i, ": ", p)
			this_out := filepath.Join(query_out, fmt.Sprint(i)+"_"+fmt.Sprint(p.Id)+"."+p.File.Ext)
			// download
			dl.DL(p.File.Url, this_out)
		}
	},
}

func init() {
	queryCmd.Flags().StringSliceVar(&query_tags, "tags", []string{}, "Tags to searh for, comma separated. i.e. male,female")
	queryCmd.Flags().IntVar(&query_first, "first", 1, "How many posts to download")
	queryCmd.Flags().StringVar(&query_out, "out", "output", "Output directory")
}
