package dl

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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
}
