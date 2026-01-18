package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dghubble/oauth1"
)

type Config struct {
	APIKey            string
	APISecret         string
	AccessToken       string
	AccessTokenSecret string
}

type Client struct {
	httpClient *http.Client
}

func NewClient(cfg Config) *Client {
	config := oauth1.NewConfig(cfg.APIKey, cfg.APISecret)
	token := oauth1.NewToken(cfg.AccessToken, cfg.AccessTokenSecret)
	return &Client{
		httpClient: config.Client(oauth1.NoContext, token),
	}
}

func (c *Client) Post(text string) error {
	return c.post(TweetRequest{Text: text})
}

func (c *Client) PostWithImage(text string, imagePath string) error {
	mediaID, err := c.uploadMedia(imagePath)
	if err != nil {
		return fmt.Errorf("upload media: %w", err)
	}

	return c.post(TweetRequest{
		Text:  text,
		Media: &MediaIDs{MediaIDs: []string{mediaID}},
	})
}

func (c *Client) post(tweet TweetRequest) error {
	body, err := json.Marshal(tweet)
	if err != nil {
		return fmt.Errorf("marshal tweet: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, tweetEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("twitter API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	return nil
}

func (c *Client) uploadMedia(imagePath string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("open image: %w", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("media", filepath.Base(imagePath))
	if err != nil {
		return "", fmt.Errorf("create form file: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return "", fmt.Errorf("copy file: %w", err)
	}
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, mediaUploadEndpoint, &buf)
	if err != nil {
		return "", fmt.Errorf("create upload request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("send upload request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("media upload error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var mediaResp struct {
		MediaID       int64  `json:"media_id"`
		MediaIDString string `json:"media_id_string"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&mediaResp); err != nil {
		return "", fmt.Errorf("decode media response: %w", err)
	}

	return mediaResp.MediaIDString, nil
}
