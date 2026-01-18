package domain

type Quote struct {
	Text      string
	Character string
	Episode   string
}

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
	Quotes   []Quote
	Images   []Image
	Opinions map[string][]string
}

func (c *Content) Total() int {
	return len(c.Quotes) + len(c.Images)
}

func (c *Content) IsEmpty() bool {
	return c.Total() == 0
}
