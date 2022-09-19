package e6api

// https://e621.net/help/api

type File struct {
	Width  int
	Height int
	Ext    string
	Size   int
	Md5    string
	Url    string
}

type Preview struct {
	Width  int
	Height int
	Url    string
}

type Sample struct {
	Has    bool
	Width  int
	Height int
	Url    string
}

type Score struct {
	Up    int
	Down  int
	Total int
}

type Tags struct {
	General   []string
	Species   []string
	Character []string
	Artist    []string
	Invalid   []string
	Lore      []string
	Meta      []string
}

type Flags struct {
	Pending       bool
	Flagged       bool
	Note_locked   bool
	Status_locked bool
	Rating_locked bool
	Deleted       bool
}

type Relationships struct {
	Parent_id           int
	Has_children        bool
	Has_active_children bool
	Children            []int
}

type Post struct {
	Id            int
	Created_at    string
	Updated_at    string
	File          File
	Preview       Preview
	Sample        Sample
	Score         Score
	Locked_tags   []string
	Change_seq    int
	Flags         Flags
	Rating        string
	Fav_count     int
	Sources       []string
	Pools         []int
	Relationships Relationships
	Approver_id   int
	Uploader_id   int
	Description   string
	Comment_count int
	Is_favorited  bool
}
