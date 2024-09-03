package smileidentity

import (
	"context"
	"net/http"
)

func (c *Client) EnhancedKYCAsyncVerification(ctx context.Context, input *KYCInput) (*AsyncResponse, error) {
	var resp AsyncResponse

	err := c.makeRequest(ctx, http.MethodPost, "v1/async_id_verification", nil, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) EnhancedKYCVerification(ctx context.Context, input *KYCInput) (*KYCVerificationResult, error) {
	var resp KYCVerificationResult

	err := c.makeRequest(ctx, http.MethodPost, "v1/id_verification", nil, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
