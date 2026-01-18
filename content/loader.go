package content

import "github.com/ray-q/umineko_bot/domain"

type Loader interface {
	Load() (*domain.Content, error)
}
