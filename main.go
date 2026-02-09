package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/ray-q/umineko_bot/api"
	"github.com/ray-q/umineko_bot/api/twitter"
	"github.com/ray-q/umineko_bot/bot"
	"github.com/ray-q/umineko_bot/content"
	"github.com/ray-q/umineko_bot/content/loader"
	"github.com/ray-q/umineko_bot/content/picker"
	"github.com/ray-q/umineko_bot/quote"
)

func main() {
	dryRun := flag.Bool("dry-run", false, "Print what would be posted without actually posting")
	mode := ModeRandom
	flag.Var(&mode, "mode", "Post mode: random, erika")
	flag.Parse()

	_ = godotenv.Load()

	dataDir := resolveDataDir()
	log.Printf("Using data directory: %s", dataDir)

	var poster api.Poster
	if *dryRun {
		poster = &api.DryRunPoster{}
	} else {
		poster = twitter.NewClient(twitter.Config{
			APIKey:            os.Getenv("TWITTER_API_KEY"),
			APISecret:         os.Getenv("TWITTER_API_SECRET"),
			AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
			AccessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		})
	}

	contentLoader := loader.NewFileLoader(dataDir)
	statePath := resolveStatePath()
	log.Printf("Using state file: %s", statePath)

	quoteURL := os.Getenv("QUOTE_SERVICE_URL")
	if quoteURL == "" {
		quoteURL = "https://quotes.auaurora.moe"
	}
	quoteClient := quote.NewClient(quoteURL)

	var p content.Picker
	switch mode {
	case ModeErika:
		p = picker.NewErikaPicker(dataDir, statePath, quoteClient)
	case ModeRandom:
		p = picker.NewRandomPicker(dataDir, statePath, quoteClient)
	}

	b := bot.New(poster, contentLoader, p)

	if err := b.Run(); err != nil {
		log.Fatal(err)
	}

	if *dryRun {
		log.Println("Dry run complete!")
	} else {
		log.Println("Successfully posted!")
	}
}

func resolveDataDir() string {
	if dir := os.Getenv("DATA_DIR"); dir != "" {
		return dir
	}

	if execPath, err := os.Executable(); err == nil {
		dir := filepath.Join(filepath.Dir(execPath), "data")
		if _, err := os.Stat(dir); err == nil {
			return dir
		}
	}

	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "data")
}

func resolveStatePath() string {
	if path := os.Getenv("STATE_PATH"); path != "" {
		return path
	}
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, ".state", "state.json")
}
