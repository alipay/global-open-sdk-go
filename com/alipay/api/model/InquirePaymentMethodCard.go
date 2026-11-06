package model

type InquirePaymentMethodCard struct {
	CardToken      string    `json:"cardToken,omitempty"`
	CardNo         string    `json:"cardNo,omitempty"`
	Mask           string    `json:"mask,omitempty"`
	Brand          string    `json:"brand,omitempty"`
	CountryIssue   string    `json:"countryIssue,omitempty"`
	CardIssuer     string    `json:"cardIssuer,omitempty"`
	ExpiryMonth    string    `json:"expiryMonth,omitempty"`
	ExpiryYear     string    `json:"expiryYear,omitempty"`
	InstUserName   *UserName `json:"instUserName,omitempty"`
	BillingAddress *Address  `json:"billingAddress,omitempty"`
}
