package model

type Card struct {
	CardNo         string    `json:"cardNo,omitempty"`
	Cvv            string    `json:"cvv,omitempty"`
	ExpiryYear     string    `json:"expiryYear,omitempty"`
	ExpiryMonth    string    `json:"expiryMonth,omitempty"`
	CardholderName *UserName `json:"cardholderName,omitempty"`
}
