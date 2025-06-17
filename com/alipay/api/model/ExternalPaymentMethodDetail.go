package model

type ExternalPaymentMethodDetail struct {
	AssetToken                  string `json:"assetToken,omitempty"`
	AccountDisplayName          string `json:"accountDisplayName,omitempty"`
	DisableReason               string `json:"disableReason,omitempty"`
	PaymentMethodDetailMetadata string `json:"paymentMethodDetailMetadata,omitempty"`
}
