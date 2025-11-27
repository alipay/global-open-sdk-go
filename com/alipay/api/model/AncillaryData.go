package model

type AncillaryData struct {
	Services              []*Service `json:"services,omitempty"`
	ConnectedTicketNumber string     `json:"connectedTicketNumber,omitempty"`
}
