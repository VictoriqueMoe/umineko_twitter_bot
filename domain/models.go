package domain

type Image struct {
	Path        string
	Description string
	Character   string
}

type Opinion struct {
	Character string   `json:"character"`
	Opinions  []string `json:"opinions"`
}

type Post struct {
	Text      string
	ImagePath string
}

type Content struct {
	Opinions map[string][]string
}
