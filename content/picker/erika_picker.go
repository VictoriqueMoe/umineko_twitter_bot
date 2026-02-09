package picker

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ray-q/umineko_bot/domain"
	"github.com/ray-q/umineko_bot/quote"
	"github.com/ray-q/umineko_bot/state"
)

type ErikaPicker struct {
	dataDir     string
	statePath   string
	quoteClient *quote.Client
}

func NewErikaPicker(dataDir, statePath string, quoteClient *quote.Client) *ErikaPicker {
	return &ErikaPicker{dataDir: dataDir, statePath: statePath, quoteClient: quoteClient}
}

func (p *ErikaPicker) Pick(content *domain.Content) domain.Post {
	erikaDir := filepath.Join(p.dataDir, "images", "erika")

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

	imagePath, relPath, err := pickRandomFile(erikaDir, s.ErikaHistory)
	if err != nil {
		log.Printf("Error picking erika image: %v", err)
		return domain.Post{Text: "No Erika images available."}
	}

	if p.statePath != "" {
		s.AddErikaPost(relPath)
		if err := s.Save(p.statePath); err != nil {
			log.Printf("Warning: could not save state: %v", err)
		}
	}

	text := ""
	if p.quoteClient != nil {
		q := p.quoteClient.RandomQuote(quote.Erika)
		if q != "" {
			text = fmt.Sprintf("「%s」", q)
		}
	}

	return domain.Post{Text: text, ImagePath: imagePath}
}
