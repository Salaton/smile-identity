package smileidentity

import (
	"context"
	"net/http"
	"time"
)

func (c *Client) VerifyPhoneNumberAsync(ctx context.Context, input *PhoneNumberVerification) (*AsyncResponse, error) {
	var resp AsyncResponse

	headers := c.phoneVerificationRequestHeaders()

	err := c.makeRequest(ctx, http.MethodPost, "v2/async-verify-phone", nil, headers, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) VerifyPhoneNumber(
	ctx context.Context,
	input *PhoneNumberVerification,
) (*PhoneNumberVerificationResponse, error) {
	var resp PhoneNumberVerificationResponse

	headers := c.phoneVerificationRequestHeaders()

	err := c.makeRequest(ctx, http.MethodPost, "v2/verify-phone-number", nil, headers, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) phoneVerificationRequestHeaders() http.Header {
	header := http.Header{}

	header.Add("smileid-partner-id", c.partnerID)
	header.Add("smileid-request-signature", c.generateSignature())
	header.Add("smileid-timestamp", time.Now().Format(time.RFC3339))
	header.Add("smileid-source-sdk", "rest_api")
	header.Add("smileid-source-sdk-version", "1.0.0")

	return header
}
