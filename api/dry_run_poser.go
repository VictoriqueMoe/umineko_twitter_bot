package api

import "fmt"

type DryRunPoster struct{}

func (d *DryRunPoster) Post(text string) error {
	fmt.Println("=== DRY RUN - Would post: ===")
	fmt.Println(text)
	fmt.Println("=============================")
	return nil
}

func (d *DryRunPoster) PostWithImage(text string, imagePath string) error {
	fmt.Println("=== DRY RUN - Would post: ===")
	fmt.Println(text)
	fmt.Printf("Image: %s\n", imagePath)
	fmt.Println("=============================")
	return nil
}
