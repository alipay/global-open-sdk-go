package model

type Merchant struct {
	ReferenceMerchantId  string       `json:"referenceMerchantId,omitempty"`
	MerchantMCC          string       `json:"merchantMCC,omitempty"`
	MerchantName         string       `json:"merchantName,omitempty"`
	MerchantDisplayName  string       `json:"merchantDisplayName,omitempty"`
	MerchantAddress      *Address     `json:"merchantAddress,omitempty"`
	MerchantRegisterDate string       `json:"merchantRegisterDate,omitempty"`
	Store                *Store       `json:"store,omitempty"`
	MerchantType         MerchantType `json:"merchantType,omitempty"`
}
