package model

type CardVerificationResult struct {
	AuthenticationType   string             `json:"authenticationType,omitempty"`
	AuthenticationMethod string             `json:"authenticationMethod,omitempty"`
	CvvResult            string             `json:"cvvResult,omitempty"`
	AvsResult            string             `json:"avsResult,omitempty"`
	AuthorizationCode    string             `json:"authorizationCode,omitempty"`
	ThreeDSResult        *RiskThreeDSResult `json:"threeDSResult,omitempty"`
}
