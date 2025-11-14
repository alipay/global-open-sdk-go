package model

type Gaming struct {
	GameName        string `json:"gameName,omitempty"`
	ToppedUpUser    string `json:"toppedUpUser,omitempty"`
	ToppedUpEmail   string `json:"toppedUpEmail,omitempty"`
	ToppedUpPhoneNo string `json:"toppedUpPhoneNo,omitempty"`
}
