package model

type Passenger struct {
	PassengerName    *UserName `json:"passengerName,omitempty"`
	PassengerEmail   string    `json:"passengerEmail,omitempty"`
	PassengerPhoneNo string    `json:"passengerPhoneNo,omitempty"`
}
