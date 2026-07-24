package model

type Customer struct {
	CustomerId        string `json:"customerId,omitempty"`
	CustomerRequestId string `json:"customerRequestId,omitempty"`
	Email             string `json:"email,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	Status            string `json:"status,omitempty"`
}
