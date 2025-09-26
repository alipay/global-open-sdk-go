package model

type Individual struct {
	Name         *UserName      `json:"name,omitempty"`
	EnglishName  *UserName      `json:"englishName,omitempty"`
	DateOfBirth  string         `json:"dateOfBirth,omitempty"`
	PlaceOfBirth *Address       `json:"placeOfBirth,omitempty"`
	Certificates []*Certificate `json:"certificates,omitempty"`
	Nationality  string         `json:"nationality,omitempty"`
	Contacts     []*Contact     `json:"contacts,omitempty"`
}
