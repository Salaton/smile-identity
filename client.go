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

type Config struct {
	APIKey      string `json:"apiKey,omitempty"`
	PartnerID   string `json:"partnerID,omitempty"`
	BaseURL     string `json:"baseURL,omitempty"`
	CallbackURL string `json:"callbackURL,omitempty"`

	HTTPClient *http.Client `json:"httpClient,omitempty"`
}

// NewClient creates a new smile id api client.
func NewClient(config Config) (*Client, error) {
	switch {
	case config.APIKey == "":
		return nil, errors.New("API key is empty")
	case config.PartnerID == "":
		return nil, errors.New("partner ID is empty")
	case config.BaseURL == "":
		return nil, errors.New("baseURL is empty")
	}

	client := &Client{
		apiKey:      config.APIKey,
		partnerID:   config.PartnerID,
		HTTP:        config.HTTPClient,
		callBackURL: config.CallbackURL,
		baseURL:     config.BaseURL,
	}

	if client.HTTP == nil {
		client.HTTP = &http.Client{}
	}

	return client, nil
}

// NewClientFromEnvVars creates a new client where the needed fields are
// retrieved from the environment variables.
func NewClientFromEnvVars() (*Client, error) {
	return NewClient(
		Config{
			APIKey:      os.Getenv("SMILE_ID_API_KEY"),
			PartnerID:   os.Getenv("SMILE_ID_PARTNER_ID"),
			BaseURL:     os.Getenv("SMILE_ID_BASE_URL"),
			CallbackURL: os.Getenv("SMILE_ID_CALLBACK_URL"),
			HTTPClient:  &http.Client{},
		},
	)
}

// generateSignature generates a base64-encoded HMAC-SHA256 signature using a timestamp,
// partner ID, and a fixed string ("sid_request"),
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
