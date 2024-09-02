package smileidentity

import (
	"context"
	"net/http"
	"time"
)

type KYCInput struct {
	CallbackURL      string        `json:"callback_url"`
	Country          string        `json:"country"`
	DOB              time.Time     `json:"dob"`
	FirstName        string        `json:"first_name"`
	Gender           string        `json:"gender"`
	IDNumber         string        `json:"id_number"`
	IDType           string        `json:"id_type"`
	LastName         string        `json:"last_name"`
	MiddleName       string        `json:"middle_name"`
	PartnerID        string        `json:"partner_id"`
	PartnerParams    PartnerParams `json:"partner_params"`
	PhoneNumber      string        `json:"phone_number"`
	Signature        string        `json:"signature"`
	SourceSDKVersion string        `json:"source_sdk_version"`
	SourceSDK        string        `json:"source_sdk"`
	Timestamp        time.Time     `json:"timestamp"`
}

type PartnerParams struct {
	JobID   string `json:"job_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
	JobType int    `json:"job_type,omitempty"`
}

type Actions struct {
	DOB                MatchResult            `json:"DOB,omitempty"`
	Gender             GenderMatch            `json:"Gender,omitempty"`
	IDVerification     MatchResult            `json:"ID_Verification,omitempty"`
	Names              MatchResult            `json:"Names,omitempty"`
	PhoneNumber        MatchResult            `json:"Phone_Number,omitempty"`
	ReturnPersonalInfo ReturnPersonalInfoEnum `json:"Return_Personal_Info,omitempty"`
	VerifyIDNumber     VerifyIDNumber         `json:"Verify_ID_Number,omitempty"`
}

type BasicKYCVerificationResult struct {
	SmileJobID    string        `json:"SmileJobID,omitempty"`
	PartnerParams PartnerParams `json:"PartnerParams,omitempty"`
	ResultText    ResultText    `json:"ResultText,omitempty"`
	ResultCode    string        `json:"ResultCode,omitempty"`
	Actions       Actions       `json:"Actions,omitempty"`
	Source        string        `json:"Source,omitempty"`
	Signature     string        `json:"signature,omitempty"`
	Timestamp     time.Time     `json:"timestamp,omitempty"`
}

func (c *Client) BasicKYCAsyncVerification(ctx context.Context, input *KYCInput) (*BasicKYCVerificationResult, error) {
	var resp BasicKYCVerificationResult

	err := c.makeRequest(ctx, http.MethodPost, "v2/verify_async", nil, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) BasicKYCVerification(ctx context.Context, input *KYCInput) (*BasicKYCVerificationResult, error) {
	var resp BasicKYCVerificationResult

	err := c.makeRequest(ctx, http.MethodPost, "v2/verify", nil, input, resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
