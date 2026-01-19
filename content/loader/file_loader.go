package loader

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

	if err := l.loadOpinions(content); err != nil {
		return nil, err
	}

	return content, nil
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
