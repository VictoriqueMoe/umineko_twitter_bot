package quote

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type quoteResponse struct {
	Text string `json:"text"`
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) RandomQuote(char Character) string {
	url := fmt.Sprintf("%s/api/v1/random?character=%s&lang=en", c.baseURL, string(char))
	resp, err := c.httpClient.Get(url)
	if err != nil {
		log.Printf("Warning: failed to fetch quote: %v", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Warning: failed to read quote response: %v", err)
		return ""
	}

	var q quoteResponse
	if err := json.Unmarshal(body, &q); err != nil {
		log.Printf("Warning: failed to parse quote response: %v", err)
		return ""
	}

	return q.Text
}
