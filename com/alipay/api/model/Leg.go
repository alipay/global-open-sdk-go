package model

type Leg struct {
	DepartureTime        string    `json:"departureTime,omitempty"`
	ArrivalTime          string    `json:"arrivalTime,omitempty"`
	DepartureAddress     *Address  `json:"departureAddress,omitempty"`
	ArrivalAddress       *Address  `json:"arrivalAddress,omitempty"`
	CarrierName          string    `json:"carrierName,omitempty"`
	CarrierNo            string    `json:"carrierNo,omitempty"`
	ClassType            ClassType `json:"classType,omitempty"`
	DepartureAirportCode string    `json:"departureAirportCode,omitempty"`
	ArrivalAirportCode   string    `json:"arrivalAirportCode,omitempty"`
	FareBasis            string    `json:"fareBasis,omitempty"`
	CouponNumber         string    `json:"couponNumber,omitempty"`
	FlightNumber         string    `json:"flightNumber,omitempty"`
}
