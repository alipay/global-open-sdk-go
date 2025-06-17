package model

type Buyer struct {
	ReferenceBuyerId      string    `json:"referenceBuyerId,omitempty"`
	BuyerName             *UserName `json:"buyerName,omitempty"`
	BuyerPhoneNo          string    `json:"buyerPhoneNo,omitempty"`
	BuyerEmail            string    `json:"buyerEmail,omitempty"`
	BuyerRegistrationTime string    `json:"buyerRegistrationTime,omitempty"`
	IsAccountVerified     bool      `json:"isAccountVerified,omitempty"`
	SuccessfulOrderCount  int32     `json:"successfulOrderCount,omitempty"`
}
