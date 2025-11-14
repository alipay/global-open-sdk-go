package model

type CreditPayPlan struct {
	InstallmentNum   int32            `json:"installmentNum,omitempty"`
	Interval         string           `json:"interval,omitempty"`
	CreditPayFeeType CreditPayFeeType `json:"creditPayFeeType,omitempty"`
	FeePercentage    int32            `json:"feePercentage,omitempty"`
}
