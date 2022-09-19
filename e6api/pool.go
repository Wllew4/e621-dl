package e6api

type Pool struct {
	Id           int
	Name         string
	Created_at   string
	Updated_at   string
	Creator_id   int
	Description  string
	Is_active    bool
	Category     string
	Post_ids     []int
	Creator_name string
	Post_count   int
}
