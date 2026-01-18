package twitter

const (
	tweetEndpoint       = "https://api.x.com/2/tweets"
	mediaUploadEndpoint = "https://upload.twitter.com/1.1/media/upload.json"
)

type TweetRequest struct {
	Text  string    `json:"text"`
	Media *MediaIDs `json:"media,omitempty"`
}

type MediaIDs struct {
	MediaIDs []string `json:"media_ids"`
}
