package smileidentity

type ResultText string

const (
	PartialMatch ResultText = "Partial Match"
	ExactMatch   ResultText = "Exact Match"
	NoMatch      ResultText = "No Match"
)

type VerifyIDNumber string

const (
	Verified          VerifyIDNumber = "Verified"
	NotVerified       VerifyIDNumber = "Not Verified"
	NotDone           VerifyIDNumber = "Not Done"
	IssuerUnavailable VerifyIDNumber = "Issuer Unavailable"
)

type MatchResult string

const (
	MatchExact      MatchResult = "Exact Match"
	MatchPartial    MatchResult = "Partial Match"
	MatchTransposed MatchResult = "Transposed"
	MatchNoMatch    MatchResult = "No Match"
)

type GenderMatch string

const (
	GenderExact       GenderMatch = "Exact Match"
	GenderNoMatch     GenderMatch = "No Match"
	GenderNotProvided GenderMatch = "Not Provided"
)

type ReturnPersonalInfoEnum string

const (
	Returned      ReturnPersonalInfoEnum = "Returned"
	NotReturned   ReturnPersonalInfoEnum = "Not Returned"
	NotApplicable ReturnPersonalInfoEnum = "Not Applicable"
)
