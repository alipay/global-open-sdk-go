package model

type AccountLastModified struct {
	PasswordChangeDate string `json:"passwordChangeDate,omitempty"`
	EmailChangeDate    string `json:"emailChangeDate,omitempty"`
	ListingChangeDate  string `json:"listingChangeDate,omitempty"`
	LoginDate          string `json:"loginDate,omitempty"`
	AddressChangeDate  string `json:"addressChangeDate,omitempty"`
}
