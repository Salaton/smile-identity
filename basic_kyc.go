package smileidentity

import (
	"context"
	"net/http"
)

func (c *Client) BasicKYCAsyncVerification(ctx context.Context, input *KYCInput) (*KYCVerificationResult, error) {
	var resp KYCVerificationResult

	err := c.makeRequest(ctx, http.MethodPost, "v2/verify_async", nil, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) BasicKYCVerification(ctx context.Context, input *KYCInput) (*KYCVerificationResult, error) {
	var resp KYCVerificationResult

	err := c.makeRequest(ctx, http.MethodPost, "v2/verify", nil, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
