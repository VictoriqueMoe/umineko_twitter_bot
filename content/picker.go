package content

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ray-q/umineko_bot/domain"
)

type RandomPicker struct {
	rng      *rand.Rand
	textOnly bool
}

func NewRandomPicker() *RandomPicker {
	return &RandomPicker{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func NewRandomPickerTextOnly() *RandomPicker {
	return &RandomPicker{
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		textOnly: true,
	}
}

func (p *RandomPicker) Pick(content *domain.Content) domain.Post {
	totalQuotes := len(content.Quotes)
	totalImages := len(content.Images)

	if p.textOnly || totalImages == 0 {
		if totalQuotes == 0 {
			return domain.Post{Text: "When the seagulls cry, none shall be left alive."}
		}
		idx := p.rng.Intn(totalQuotes)
		return p.formatQuote(content.Quotes[idx])
	}

	total := totalQuotes + totalImages
	if total == 0 {
		return domain.Post{Text: "When the seagulls cry, none shall be left alive."}
	}

	idx := p.rng.Intn(total)

	if idx < totalQuotes {
		return p.formatQuote(content.Quotes[idx])
	}

	return p.formatImage(content.Images[idx-totalQuotes], content)
}

func (p *RandomPicker) formatQuote(quote domain.Quote) domain.Post {
	text := fmt.Sprintf("\"%s\"\n\n— %s", quote.Text, quote.Character)
	if quote.Episode != "" {
		text += fmt.Sprintf(" (%s)", quote.Episode)
	}
	return domain.Post{Text: text}
}

func (p *RandomPicker) formatImage(image domain.Image, content *domain.Content) domain.Post {
	var parts []string

	if image.Character != "" {
		charKey := strings.ToLower(image.Character)
		if opinions, ok := content.Opinions[charKey]; ok && len(opinions) > 0 {
			opinion := opinions[p.rng.Intn(len(opinions))]
			parts = append(parts, opinion)
		}
	}

	var hashtag string
	if image.Character != "" {
		hashtag = fmt.Sprintf("#%s #UminekoNoNakuKoroNi", strings.ReplaceAll(image.Character, " ", ""))
	} else {
		hashtag = "#UminekoNoNakuKoroNi"
	}
	parts = append(parts, hashtag)

	return domain.Post{Text: strings.Join(parts, "\n\n"), ImagePath: image.Path}
}

type ErikaPicker struct {
	rng *rand.Rand
}

func NewErikaPicker() *ErikaPicker {
	return &ErikaPicker{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
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

	idx := p.rng.Intn(len(erikaImages))
	return domain.Post{ImagePath: erikaImages[idx].Path}
}
