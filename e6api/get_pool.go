package e6api

import "fmt"

type poolResponse []Pool

func GetPoolById(id int) Pool {
	return getPool(fmt.Sprint(id))
}

func GetPoolByUrl(url string) Pool {
	id := ParseId(url)
	return getPool(id)
}

func getPool(id string) Pool {
	var unmarshalled poolResponse
	fetch_json("https://e621.net/pools.json?search[id]="+id, &unmarshalled)
	return unmarshalled[0]
}
