package model

type PassengerIdType string

const (
	PassengerIdType_PASSPORT                PassengerIdType = "PASSPORT"
	PassengerIdType_NATIONAL_ID_CARD        PassengerIdType = "NATIONAL_ID_CARD"
	PassengerIdType_DRIVER_LICENSE          PassengerIdType = "DRIVER_LICENSE"
	PassengerIdType_MILITARY_ID             PassengerIdType = "MILITARY_ID"
	PassengerIdType_GREEN_CARD              PassengerIdType = "GREEN_CARD"
	PassengerIdType_TRAVEL_DOCUMENT         PassengerIdType = "TRAVEL_DOCUMENT"
	PassengerIdType_ALIEN_REGISTRATION_CARD PassengerIdType = "ALIEN_REGISTRATION_CARD"
	PassengerIdType_BIRTH_CERTIFICATE       PassengerIdType = "BIRTH_CERTIFICATE"
	PassengerIdType_OTHERS                  PassengerIdType = "OTHERS"
)
