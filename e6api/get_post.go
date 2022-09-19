package e6api

import "fmt"

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

func getPost(id string) Post {
	var unmarshalled postResponse
	fetch_json("https://e621.net/posts.json?tags=id:"+id, &unmarshalled)
	return unmarshalled.Posts[0]
}
