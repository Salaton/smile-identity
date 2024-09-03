package smileidentity

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

func (c *Client) VerifyPhoneNumberAsync(ctx context.Context, input *PhoneNumberVerification) (*AsyncResponse, error) {
	var resp AsyncResponse

	params := c.phoneVerificationRequestHeaders()

	err := c.makeRequest(ctx, http.MethodPost, "v2/async-verify-phone", params, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) VerifyPhoneNumber(ctx context.Context, input *PhoneNumberVerification) (*PhoneNumberVerificationResponse, error) {
	var resp PhoneNumberVerificationResponse

	params := c.phoneVerificationRequestHeaders()

	err := c.makeRequest(ctx, http.MethodPost, "v2/verify-phone-number", params, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) phoneVerificationRequestHeaders() url.Values {
	params := url.Values{}
	params.Add("smileid-partner-id", c.partnerID)
	params.Add("smileid-request-signature", c.generateSignature())
	params.Add("smileid-timestamp", time.Now().String())
	params.Add("smileid-source-sdk", "rest_api")
	params.Add("smileid-source-sdk-version", "1.0.0")

	return params
}
