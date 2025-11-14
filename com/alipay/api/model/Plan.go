package model

type Plan struct {
	InterestRate         string  `json:"interestRate,omitempty"`
	MinInstallmentAmount *Amount `json:"minInstallmentAmount,omitempty"`
	MaxInstallmentAmount *Amount `json:"maxInstallmentAmount,omitempty"`
	InstallmentNum       string  `json:"installmentNum,omitempty"`
	Interval             string  `json:"interval,omitempty"`
	Enabled              bool    `json:"enabled,omitempty"`
}
