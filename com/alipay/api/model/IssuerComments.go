package model

// IssuerComments contains issuer and cardholder comments returned in a dispute notification.
type IssuerComments struct {
	CardholderComments           string `json:"cardholderComments,omitempty"`
	ReasonOfInvalidAuthorization string `json:"reasonOfInvalidAuthorization,omitempty"`
	ExplanationOfCreditPresented string `json:"explanationOfCreditPresented,omitempty"`
	JudgeReason                  string `json:"judgeReason,omitempty"`
}
