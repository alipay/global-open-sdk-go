package model

type Transit struct {
	TransitType TransitType  `json:"transitType,omitempty"`
	Legs        []*Leg       `json:"legs,omitempty"`
	Passengers  []*Passenger `json:"passengers,omitempty"`
}
