package smileidentity

import "time"

type KYCInput struct {
	CallbackURL      string        `json:"callback_url,omitempty"`
	Country          string        `json:"country,omitempty"`
	DOB              time.Time     `json:"dob,omitempty"`
	FirstName        string        `json:"first_name,omitempty"`
	Gender           string        `json:"gender,omitempty"`
	IDNumber         string        `json:"id_number,omitempty"`
	IDType           string        `json:"id_type,omitempty"`
	LastName         string        `json:"last_name,omitempty"`
	MiddleName       string        `json:"middle_name,omitempty"`
	PartnerID        string        `json:"partner_id,omitempty"`
	PartnerParams    PartnerParams `json:"partner_params,omitempty"`
	PhoneNumber      string        `json:"phone_number,omitempty"`
	Signature        string        `json:"signature,omitempty"`
	SourceSDKVersion string        `json:"source_sdk_version,omitempty"`
	SourceSDK        string        `json:"source_sdk,omitempty"`
	Timestamp        time.Time     `json:"timestamp,omitempty"`
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

type KYCVerificationResult struct {
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
	Source            string        `json:"Source,omitempty"`
}

type AsyncResponse struct {
	Success bool `json:"success,omitempty"`
}
