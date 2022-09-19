package dl

import (
	"e621dl/e6api"
	"fmt"
	"os"
	"path/filepath"
)

func DLPool(pool e6api.Pool, dirname string) {
	err := os.MkdirAll(pool.Name, os.ModePerm)
	if err != nil {
		panic(err)
	}
	for i, id := range pool.Post_ids {
		post := e6api.GetPostById(id)
		DL(post.File.Url, filepath.Join(dirname, fmt.Sprint(i)+"_"+fmt.Sprint(post.Id)+"."+post.File.Ext))
	}
}
