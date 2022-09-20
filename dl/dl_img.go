package dl

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var total_downloaded = 0

func incDownloaded() {
	total_downloaded++
}

func GetDownloadTotal() int {
	return total_downloaded
}

func DL(url string, path string) {
	fmt.Println("Downloading: " + url)
	f, _ := os.Create(path)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done: " + path)
	incDownloaded()
}
