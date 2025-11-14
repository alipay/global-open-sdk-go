package model

type Lodging struct {
	HotelName      string      `json:"hotelName,omitempty"`
	HotelAddress   *Address    `json:"hotelAddress,omitempty"`
	CheckInDate    string      `json:"checkInDate,omitempty"`
	CheckOutDate   string      `json:"checkOutDate,omitempty"`
	NumberOfNights int32       `json:"numberOfNights,omitempty"`
	NumberOfRooms  int32       `json:"numberOfRooms,omitempty"`
	GuestNames     []*UserName `json:"guestNames,omitempty"`
}
