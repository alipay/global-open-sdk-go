package model

type Passenger struct {
	PassengerName    *UserName       `json:"passengerName,omitempty"`
	PassengerEmail   string          `json:"passengerEmail,omitempty"`
	PassengerPhoneNo string          `json:"passengerPhoneNo,omitempty"`
	PassengerId      string          `json:"passengerId,omitempty"`
	PassengerIdType  PassengerIdType `json:"passengerIdType,omitempty"`
	PassengerCode    string          `json:"passengerCode,omitempty"`
}
