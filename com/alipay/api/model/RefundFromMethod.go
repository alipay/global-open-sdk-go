package model

type RefundFromMethod struct {
	GrantToken           string `json:"grantToken,omitempty"`
	RefundFromMethodType string `json:"refundFromMethodType,omitempty"`
	CustomerId           string `json:"customerId,omitempty"`
}
