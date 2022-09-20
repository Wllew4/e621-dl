package e6api

import (
	"fmt"
	"strings"
)

type postResponse struct {
	Posts []Post
}

func GetPostById(id int) Post {
	return getPost(fmt.Sprint(id))
}

func GetPostByUrl(url string) Post {
	id := ParseId(url)
	return getPost(id)
}

func GetPostsByTags(tags []string, first int) []Post {
	var unmarshalled postResponse
	tags_query := strings.Join(tags, " ")
	fetch_json("https://e621.net/posts.json"+
		"?tags="+tags_query+
		"&limit="+fmt.Sprint(first),
		&unmarshalled)
	return unmarshalled.Posts
}

func getPost(id string) Post {
	return GetPostsByTags([]string{"id:" + id}, 1)[0]
}
