package model

type Transit struct {
	TransitType               TransitType    `json:"transitType,omitempty"`
	Legs                      []*Leg         `json:"legs,omitempty"`
	Passengers                []*Passenger   `json:"passengers,omitempty"`
	AgentCode                 string         `json:"agentCode,omitempty"`
	AgentName                 string         `json:"agentName,omitempty"`
	TicketNumber              string         `json:"ticketNumber,omitempty"`
	TicketIssuerCode          string         `json:"ticketIssuerCode,omitempty"`
	RestrictedTicketIndicator string         `json:"restrictedTicketIndicator,omitempty"`
	AncillaryData             *AncillaryData `json:"ancillaryData,omitempty"`
}
