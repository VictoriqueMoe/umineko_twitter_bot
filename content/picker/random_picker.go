package picker

import (
	"fmt"
	"log"
	"math/rand/v2"
	"path/filepath"
	"strings"

	"github.com/ray-q/umineko_bot/domain"
	"github.com/ray-q/umineko_bot/quote"
	"github.com/ray-q/umineko_bot/state"
)

var doubleErikaMemes = []string{
	"DOUBLE ERIKA DAY 🎉\n\nThe detective's authority is DOUBLE absolute today!",
	"DOUBLE ERIKA DAY 🎉\n\nTwo Erikas in one day? Intellectual rapture x2!",
	"DOUBLE ERIKA DAY 🎉\n\nBern blessed us with maximum Erika content today.",
	"DOUBLE ERIKA DAY 🎉\n\nToday's truth: You can never have too much Erika.",
	"DOUBLE ERIKA DAY 🎉\n\nKnox's 11th: Thou shalt celebrate double Erika days.",
	"DOUBLE ERIKA DAY 🎉\n\nThis is what peak detective content looks like.",
	"DOUBLE ERIKA DAY 🎉\n\nErika in the morning, Erika in the evening. As god intended.",
}

type RandomPicker struct {
	dataDir     string
	statePath   string
	quoteClient *quote.Client
}

const twitterCharLimit = 280

func NewRandomPicker(dataDir, statePath string, quoteClient *quote.Client) *RandomPicker {
	return &RandomPicker{dataDir: dataDir, statePath: statePath, quoteClient: quoteClient}
}

func (p *RandomPicker) Pick(content *domain.Content) domain.Post {
	imagesDir := filepath.Join(p.dataDir, "images")

	var s *state.State
	if p.statePath != "" {
		var err error
		s, err = state.Load(p.statePath)
		if err != nil {
			log.Printf("Warning: could not load state: %v", err)
			s = &state.State{}
		}
	} else {
		s = &state.State{}
	}

	imagePath, relPath, err := pickRandomFile(imagesDir, s.RandomHistory)
	if err != nil {
		log.Printf("Error picking random image: %v", err)
		return domain.Post{Text: "No images available."}
	}

	if p.statePath != "" {
		s.AddRandomPost(relPath)
		if err := s.Save(p.statePath); err != nil {
			log.Printf("Warning: could not save state: %v", err)
		}
	}

	parts := strings.Split(filepath.ToSlash(relPath), "/")
	character := ""
	if len(parts) > 0 {
		character = parts[0]
	}

	return p.formatImagePost(imagePath, character, content)
}

func (p *RandomPicker) formatImagePost(imagePath, character string, content *domain.Content) domain.Post {
	var parts []string

	isErika := strings.ToLower(character) == "erika"
	if isErika {
		parts = append(parts, doubleErikaMemes[rand.IntN(len(doubleErikaMemes))])
	}

	if character != "" && !isErika {
		charKey := strings.ToLower(character)
		if opinions, ok := content.Opinions[charKey]; ok && len(opinions) > 0 {
			opinion := opinions[rand.IntN(len(opinions))]
			parts = append(parts, opinion)
		}
	}

	var hashtag string
	if character != "" {
		cleanName := strings.TrimRight(character, "0123456789")
		if cleanName == "" {
			cleanName = character
		}
		cleanName = strings.ReplaceAll(cleanName, " ", "")
		hashtag = fmt.Sprintf("#%s #UminekoNoNakuKoroNi", cleanName)
	} else {
		hashtag = "#UminekoNoNakuKoroNi"
	}
	parts = append(parts, hashtag)

	if p.quoteClient != nil && character != "" {
		charKey := strings.ToLower(character)
		if char, ok := quote.LookupCharacter(charKey); ok {
			for i := 0; i < 3; i++ {
				q := p.quoteClient.RandomQuote(char)
				if q == "" {
					break
				}
				quoted := fmt.Sprintf("「%s」", q)
				candidate := strings.Join(append([]string{quoted}, parts...), "\n\n")
				if len(candidate) <= twitterCharLimit {
					parts = append([]string{quoted}, parts...)
					break
				}
			}
		}
	}

	return domain.Post{Text: strings.Join(parts, "\n\n"), ImagePath: imagePath}
}
