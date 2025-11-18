package model

type InterestFree struct {
	Provider            string  `json:"provider,omitempty"`
	ExpireTime          string  `json:"expireTime,omitempty"`
	InstallmentFreeNums []int32 `json:"installmentFreeNums,omitempty"`
	MinPaymentAmount    *Amount `json:"minPaymentAmount,omitempty"`
	MaxPaymentAmount    *Amount `json:"maxPaymentAmount,omitempty"`
	FreePercentage      int32   `json:"freePercentage,omitempty"`
}
