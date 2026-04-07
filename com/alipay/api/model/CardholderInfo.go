package model

type CardholderInfo struct {
	CardHolderName *UserName `json:"cardHolderName,omitempty"`
	BillAddress    *Address  `json:"billAddress,omitempty"`
	DisplayName    string    `json:"displayName,omitempty"`
}
