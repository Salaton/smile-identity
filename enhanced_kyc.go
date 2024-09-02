package smileidentity

import "time"

type EnhancedKYCVerificationResult struct {
	Actions           Actions       `json:"Actions,omitempty"`
	Country           string        `json:"Country,omitempty"`
	DOB               string        `json:"DOB,omitempty"`
	ExpirationDate    string        `json:"ExpirationDate,omitempty"`
	IssuanceDate      string        `json:"IssuanceDate,omitempty"`
	FullName          string        `json:"FullName,omitempty"`
	IDNumber          string        `json:"IDNumber,omitempty"`
	SecondaryIDNumber string        `json:"SecondaryIDNumber,omitempty"`
	IDType            string        `json:"IDType,omitempty"`
	PartnerParams     PartnerParams `json:"PartnerParams,omitempty"`
	Photo             string        `json:"Photo,omitempty"`
	ResultCode        string        `json:"ResultCode,omitempty"`
	ResultText        ResultText    `json:"ResultText,omitempty"`
	SmileJobID        string        `json:"SmileJobID,omitempty"`
	Signature         string        `json:"signature,omitempty"`
	Timestamp         time.Time     `json:"timestamp,omitempty"`
}
