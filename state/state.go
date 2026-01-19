package state

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type State struct {
	ErikaHistory  []string `json:"erika_history"`
	RandomHistory []string `json:"random_history"`
}

func Load(path string) (*State, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &State{}, nil
		}
		return nil, err
	}

	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *State) Save(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (s *State) AddErikaPost(folder string) {
	s.ErikaHistory = append(s.ErikaHistory, folder)
	if len(s.ErikaHistory) > 10 {
		s.ErikaHistory = s.ErikaHistory[len(s.ErikaHistory)-10:]
	}
}

func (s *State) AddRandomPost(path string) {
	s.RandomHistory = append(s.RandomHistory, path)
	if len(s.RandomHistory) > 10 {
		s.RandomHistory = s.RandomHistory[len(s.RandomHistory)-10:]
	}
}
