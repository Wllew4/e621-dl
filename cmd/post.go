package cmd

import (
	"e621dl/dl"
	"e621dl/e6api"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	post_url string
	post_id  int
	post_out string
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Download a post",
	Long:  "Download a post",
	Run: func(cmd *cobra.Command, args []string) {
		// verify flags
		if post_url == "" && post_id == 0 {
			fmt.Println("Provide a url or id.\ne621-dl post --url https://e621.net/posts/<id>\ne621-dl post --id <id>")
			os.Exit(1)
		}

		// fetch post data
		fmt.Println("Fetching post data")
		var post e6api.Post
		if post_id != 0 {
			post = e6api.GetPostById(post_id)
		} else {
			post = e6api.GetPostByUrl(post_url)
		}
		fmt.Println("Found: " + post.File.Url)

		// determine output file
		if post_out == "" {
			post_out = fmt.Sprint(post.Id) + "." + post.File.Ext
		}

		// download
		dl.DL(post.File.Url, post_out)
	},
}

func init() {
	postCmd.Flags().StringVar(&post_url, "url", "", "Post url, i.e. https://e621.net/posts/<id>")
	postCmd.Flags().IntVar(&post_id, "id", 0, "Post id")
	postCmd.Flags().StringVar(&post_out, "out", "", "Output file")
}
