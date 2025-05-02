package smileidentity

import (
	"context"
	"net/http"
)

func (c *Client) JobRequest(ctx context.Context, input *RequestJobPayload) (*RequestJobResponse, error) {
	var resp RequestJobResponse

	input.PartnerParams.JobType = 6

	err := c.makeRequest(ctx, http.MethodPost, "v1/upload", nil, nil, input, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) UploadJobPayload(ctx context.Context, input *RequestJobResponse) error {
	return nil
}
