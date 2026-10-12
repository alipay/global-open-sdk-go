package model

type CustomerBusinessAddress struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
	Address string `json:"address,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
}
