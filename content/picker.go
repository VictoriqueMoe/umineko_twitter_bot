package content

import "github.com/ray-q/umineko_bot/domain"

type Picker interface {
	Pick(content *domain.Content) domain.Post
}
