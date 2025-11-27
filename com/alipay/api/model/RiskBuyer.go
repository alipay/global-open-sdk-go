package model

type RiskBuyer struct {
	NoteToMerchant  string `json:"noteToMerchant,omitempty"`
	NoteToShipping  string `json:"noteToShipping,omitempty"`
	OrderCountIn1H  int32  `json:"orderCountIn1H,omitempty"`
	OrderCountIn24H int32  `json:"orderCountIn24H,omitempty"`
}
