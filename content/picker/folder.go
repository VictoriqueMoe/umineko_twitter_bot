package picker

import (
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
)

func pickRandomFile(root string, history []string) (string, string, error) {
	path, err := walkToFiles(root, history)
	if err != nil {
		return "", "", err
	}

	files, err := listImageFiles(path)
	if err != nil {
		return "", "", err
	}

	file := files[rand.IntN(len(files))]
	fullPath := filepath.Join(path, file)

	relPath, _ := filepath.Rel(root, path)
	if relPath == "." {
		relPath = ""
	}

	return fullPath, relPath, nil
}

func walkToFiles(dir string, history []string) (string, error) {
	for {
		subdirs, err := listSubdirs(dir)
		if err != nil {
			return "", err
		}

		if len(subdirs) == 0 {
			return dir, nil
		}

		hasFiles, err := hasImageFiles(dir)
		if err != nil {
			return "", err
		}

		if hasFiles && len(subdirs) > 0 {
			if rand.IntN(2) == 0 {
				return dir, nil
			}
		}

		selected := selectWithWeighting(subdirs, dir, history)
		dir = filepath.Join(dir, selected)
	}
}

func selectWithWeighting(options []string, currentPath string, history []string) string {
	if len(options) == 1 {
		return options[0]
	}

	recentCount := make(map[string]int)
	lookback := min(len(history), 5)
	for i := len(history) - lookback; i < len(history); i++ {
		histPath := history[i]
		for _, opt := range options {
			checkPath := filepath.Join(currentPath, opt)
			checkPath = filepath.ToSlash(checkPath)
			histPath = filepath.ToSlash(histPath)
			if strings.HasPrefix(histPath, checkPath) || strings.Contains(histPath, "/"+opt+"/") || strings.HasSuffix(histPath, "/"+opt) {
				recentCount[opt]++
			}
		}
	}

	weights := make(map[string]float64)
	for _, opt := range options {
		count := recentCount[opt]
		weights[opt] = 1.0 / float64(1+count*count)
	}

	var totalWeight float64
	for _, w := range weights {
		totalWeight += w
	}

	r := rand.Float64() * totalWeight
	var cumulative float64
	for _, opt := range options {
		cumulative += weights[opt]
		if r <= cumulative {
			return opt
		}
	}

	return options[0]
}

func listSubdirs(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var subdirs []string
	for _, e := range entries {
		if e.IsDir() {
			subdirs = append(subdirs, e.Name())
		}
	}
	return subdirs, nil
}

func listImageFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && isImageFile(e.Name()) {
			files = append(files, e.Name())
		}
	}
	return files, nil
}

func hasImageFiles(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}

	for _, e := range entries {
		if !e.IsDir() && isImageFile(e.Name()) {
			return true, nil
		}
	}
	return false, nil
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
