package bot

import "fmt"

type Bot struct {
	poster Poster
	loader ContentLoader
	picker Picker
}

func New(poster Poster, loader ContentLoader, picker Picker) *Bot {
	return &Bot{
		poster: poster,
		loader: loader,
		picker: picker,
	}
}

func (b *Bot) Run() error {
	content, err := b.loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load content: %w", err)
	}

	if content.IsEmpty() {
		return fmt.Errorf("no content available to post")
	}

	post := b.picker.Pick(content)

	if post.ImagePath != "" {
		return b.poster.PostWithImage(post.Text, post.ImagePath)
	}
	return b.poster.Post(post.Text)
}
