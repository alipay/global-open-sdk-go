package model

type TaxTransaction struct {
	TaxTransactionId   string `json:"taxTransactionId,omitempty"`
	TaxCalculationId   string `json:"taxCalculationId,omitempty"`
	Type               string `json:"type,omitempty"`
	TaxAmount          string `json:"taxAmount,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Status             string `json:"status,omitempty"`
	FailureReason      string `json:"failureReason,omitempty"`
	TaxDate            string `json:"taxDate,omitempty"`
	PostedAt           string `json:"postedAt,omitempty"`
	ReferencePaymentId string `json:"referencePaymentId,omitempty"`
	ReferenceRefundId  string `json:"referenceRefundId,omitempty"`
}
