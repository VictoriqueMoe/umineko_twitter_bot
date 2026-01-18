package bot

import "github.com/ray-q/umineko_bot/domain"

type Poster interface {
	Post(text string) error
	PostWithImage(text string, imagePath string) error
}

type ContentLoader interface {
	Load() (*domain.Content, error)
}

type Picker interface {
	Pick(content *domain.Content) domain.Post
}
