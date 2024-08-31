package smileidentity

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"time"
)

type Client struct {
	apiKey      string
	partnerID   string
	callBackURL string
	baseURL     string

	HTTP *http.Client
}

// NewClient creates a new smile id api client
func NewClient(apiKey, partnerID, baseURL, callbackURL string) (*Client, error) {
	switch {
	case apiKey == "":
		return nil, errors.New("API key is empty")
	case partnerID == "":
		return nil, errors.New("partner ID is empty")
	case baseURL == "":
		return nil, errors.New("baseURL is empty")
	}

	client := &Client{
		apiKey:      apiKey,
		partnerID:   partnerID,
		HTTP:        &http.Client{},
		callBackURL: callbackURL,
		baseURL:     baseURL,
	}

	return client, nil
}

// NewClientFromEnvVars creates a new client where the needed fields are
// retrieved from the environment variables
func NewClientFromEnvVars() (*Client, error) {
	return NewClient(
		os.Getenv("SMILE_ID_API_KEY"),
		os.Getenv("SMILE_ID_PARTNER_ID"),
		os.Getenv("SMILE_ID_BASE_URL"),
		os.Getenv("SMILE_ID_CALLBACK_URL"),
	)
}

// generateSignature generates a base64-encoded HMAC-SHA256 signature using a timestamp, partner ID, and a fixed string ("sid_request"),
// with the API Key as the secret key.
func (c *Client) generateSignature() string {
	timeStamp := time.Now().Format(time.RFC3339)

	h := hmac.New(sha256.New, []byte(c.apiKey))

	h.Write([]byte(timeStamp))
	h.Write([]byte(c.partnerID))
	h.Write([]byte("sid_request"))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
}
