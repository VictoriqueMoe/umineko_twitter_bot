package picker

import (
	"math/rand/v2"
	"strings"

	"github.com/ray-q/umineko_bot/domain"
)

type ErikaPicker struct{}

func NewErikaPicker() *ErikaPicker {
	return &ErikaPicker{}
}

func (p *ErikaPicker) Pick(content *domain.Content) domain.Post {
	var erikaImages []domain.Image
	for _, img := range content.Images {
		if strings.ToLower(img.Character) == "erika" {
			erikaImages = append(erikaImages, img)
		}
	}

	if len(erikaImages) == 0 {
		return domain.Post{Text: "No Erika images available."}
	}

	idx := rand.IntN(len(erikaImages))
	return domain.Post{ImagePath: erikaImages[idx].Path}
}
