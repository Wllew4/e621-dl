package e6api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetch_json(url string, out any) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &out)
	if err != nil {
		panic(err)
	}
}
