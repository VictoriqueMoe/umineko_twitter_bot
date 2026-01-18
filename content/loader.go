package content

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ray-q/umineko_bot/domain"
)

type FileLoader struct {
	dataDir string
}

func NewFileLoader(dataDir string) *FileLoader {
	return &FileLoader{dataDir: dataDir}
}

func (l *FileLoader) Load() (*domain.Content, error) {
	content := &domain.Content{
		Opinions: make(map[string][]string),
	}

	if err := l.loadQuotes(content); err != nil {
		return nil, err
	}

	if err := l.loadImages(content); err != nil {
		return nil, err
	}

	if err := l.loadOpinions(content); err != nil {
		return nil, err
	}

	return content, nil
}

func (l *FileLoader) loadQuotes(content *domain.Content) error {
	quotesPath := filepath.Join(l.dataDir, "quotes.json")
	if _, err := os.Stat(quotesPath); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(quotesPath)
	if err != nil {
		return fmt.Errorf("read quotes file: %w", err)
	}

	var quotes []domain.Quote
	if err := json.Unmarshal(data, &quotes); err != nil {
		return fmt.Errorf("parse quotes: %w", err)
	}

	content.Quotes = quotes
	return nil
}

func (l *FileLoader) loadImages(content *domain.Content) error {
	imagesDir := filepath.Join(l.dataDir, "images")
	if _, err := os.Stat(imagesDir); os.IsNotExist(err) {
		return nil
	}

	if images, err := l.loadImagesFromJSON(); err == nil {
		content.Images = images
		return nil
	}

	return l.loadImagesFromDirectory(content, imagesDir)
}

func (l *FileLoader) loadImagesFromJSON() ([]domain.Image, error) {
	imagesJsonPath := filepath.Join(l.dataDir, "images.json")
	data, err := os.ReadFile(imagesJsonPath)
	if err != nil {
		return nil, err
	}

	var images []domain.Image
	if err := json.Unmarshal(data, &images); err != nil {
		return nil, err
	}

	return images, nil
}

func (l *FileLoader) loadImagesFromDirectory(content *domain.Content, imagesDir string) error {
	return filepath.WalkDir(imagesDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		if !isImageFile(d.Name()) {
			return nil
		}

		relPath, _ := filepath.Rel(imagesDir, path)
		character := extractCharacter(relPath)

		content.Images = append(content.Images, domain.Image{
			Path:        path,
			Description: strings.TrimSuffix(d.Name(), filepath.Ext(d.Name())),
			Character:   character,
		})

		return nil
	})
}

func isImageFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif":
		return true
	default:
		return false
	}
}

func extractCharacter(relPath string) string {
	parts := strings.Split(filepath.Dir(relPath), string(filepath.Separator))
	if len(parts) > 0 && parts[0] != "." {
		return parts[0]
	}
	return ""
}

func (l *FileLoader) loadOpinions(content *domain.Content) error {
	opinionsPath := filepath.Join(l.dataDir, "opinions.json")
	if _, err := os.Stat(opinionsPath); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(opinionsPath)
	if err != nil {
		return fmt.Errorf("read opinions file: %w", err)
	}

	var opinions []domain.Opinion
	if err := json.Unmarshal(data, &opinions); err != nil {
		return fmt.Errorf("parse opinions: %w", err)
	}

	for _, op := range opinions {
		content.Opinions[strings.ToLower(op.Character)] = op.Opinions
	}

	return nil
}
