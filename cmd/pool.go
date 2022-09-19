package cmd

import (
	"e621dl/dl"
	"e621dl/e6api"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	pool_url string
	pool_id  int
	pool_out string
)

var poolCmd = &cobra.Command{
	Use:   "pool",
	Short: "Download a pool",
	Long:  "Download a pool",
	Run: func(cmd *cobra.Command, args []string) {
		// verify flags
		if pool_url == "" && pool_id == 0 {
			fmt.Println("Provide a url or id.\ne621-dl pool --url https://e621.net/pools/<id>\ne621-dl pool --id <id>")
			os.Exit(1)
		}

		// fetch pool data
		fmt.Println("Fetching pool data")
		var pool e6api.Pool
		if pool_id != 0 {
			pool = e6api.GetPoolById(pool_id)
		} else {
			pool = e6api.GetPoolByUrl(pool_url)
		}
		fmt.Println("Found: " + pool.Name)

		// determine output directory
		if pool_out == "" {
			pool_out = fmt.Sprint(pool.Name)
		}

		// download
		dl.DLPool(pool, pool_out)
	},
}

func init() {
	poolCmd.Flags().StringVar(&pool_url, "url", "", "Pool url, i.e. https://e621.net/pools/<id>")
	poolCmd.Flags().IntVar(&pool_id, "id", 0, "Pool id")
	poolCmd.Flags().StringVar(&pool_out, "out", "", "Output directory")
}
