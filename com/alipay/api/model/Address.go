package model

type Address struct {
	Region   string `json:"region,omitempty"`
	State    string `json:"state,omitempty"`
	City     string `json:"city,omitempty"`
	Address1 string `json:"address1,omitempty"`
	Address2 string `json:"address2,omitempty"`
	ZipCode  string `json:"zipCode,omitempty"`
	Label    string `json:"label,omitempty"`
}
