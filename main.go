package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/ray-q/umineko_bot/api/twitter"
	"github.com/ray-q/umineko_bot/bot"
	"github.com/ray-q/umineko_bot/content"
)

type dryRunPoster struct{}

func (d *dryRunPoster) Post(text string) error {
	fmt.Println("=== DRY RUN - Would post: ===")
	fmt.Println(text)
	fmt.Println("=============================")
	return nil
}

func (d *dryRunPoster) PostWithImage(text string, imagePath string) error {
	fmt.Println("=== DRY RUN - Would post: ===")
	fmt.Println(text)
	fmt.Printf("Image: %s\n", imagePath)
	fmt.Println("=============================")
	return nil
}

func main() {
	dryRun := flag.Bool("dry-run", false, "Print what would be posted without actually posting")
	textOnly := flag.Bool("text-only", false, "Only post text quotes, no images")
	mode := flag.String("mode", "random", "Post mode: random (character with opinion) or erika (Erika image only)")
	flag.Parse()

	_ = godotenv.Load()

	dataDir := resolveDataDir()
	log.Printf("Using data directory: %s", dataDir)

	var poster bot.Poster
	if *dryRun {
		poster = &dryRunPoster{}
	} else {
		poster = twitter.NewClient(twitter.Config{
			APIKey:            os.Getenv("TWITTER_API_KEY"),
			APISecret:         os.Getenv("TWITTER_API_SECRET"),
			AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
			AccessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		})
	}

	loader := content.NewFileLoader(dataDir)
	var picker bot.Picker
	switch *mode {
	case "erika":
		picker = content.NewErikaPicker()
	default:
		if *textOnly {
			picker = content.NewRandomPickerTextOnly()
		} else {
			picker = content.NewRandomPicker()
		}
	}

	b := bot.New(poster, loader, picker)

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
