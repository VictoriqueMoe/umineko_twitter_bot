package bot

import (
	"fmt"

	"github.com/ray-q/umineko_bot/api"
	"github.com/ray-q/umineko_bot/content"
)

type Bot struct {
	poster api.Poster
	loader content.Loader
	picker content.Picker
}

func New(poster api.Poster, loader content.Loader, picker content.Picker) *Bot {
	return &Bot{
		poster: poster,
		loader: loader,
		picker: picker,
	}
}

func (b *Bot) Run() error {
	c, err := b.loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load content: %w", err)
	}

	if c.IsEmpty() {
		return fmt.Errorf("no content available to post")
	}

	post := b.picker.Pick(c)

	if post.ImagePath != "" {
		return b.poster.PostWithImage(post.Text, post.ImagePath)
	}
	return b.poster.Post(post.Text)
}
