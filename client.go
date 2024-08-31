package smileidentity

import (
	"errors"
	"net/http"
	"os"
)

type Client struct {
	apiKey      string
	partnerID   string
	callBackURL string
	HTTP        *http.Client
}

// NewClient creates a new smile id api client
func NewClient(apiKey, partnerID, callbackURL string) (*Client, error) {
	switch {
	case apiKey == "":
		return nil, errors.New("API key is empty")
	case partnerID == "":
		return nil, errors.New("partner ID is empty")
	}

	client := &Client{
		apiKey:      apiKey,
		partnerID:   partnerID,
		HTTP:        &http.Client{},
		callBackURL: callbackURL,
	}

	return client, nil
}

// NewClientFromEnvVars creates a new client where the needed fields are
// retrieved from the environment variables
func NewClientFromEnvVars() (*Client, error) {
	return NewClient(
		os.Getenv("SMILE_ID_API_KEY"),
		os.Getenv("SMILE_ID_PARTNER_ID"),
		os.Getenv("SMILE_ID_CALLBACK_URL"),
	)
}
