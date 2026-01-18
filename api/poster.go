package api

type Poster interface {
	Post(text string) error
	PostWithImage(text string, imagePath string) error
}
