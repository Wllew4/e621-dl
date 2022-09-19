package e6api

import (
	"fmt"
	"os"
	"regexp"
)

func ParseId(url string) string {
	re := regexp.MustCompile(`.*e621.net/.*/([0-9]*)`)
	submatchall := re.FindStringSubmatch(url)
	if len(submatchall) < 2 {
		fmt.Println("Invalid url: " + url)
		os.Exit(1)
	}
	return submatchall[1]
}
