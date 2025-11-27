package model

type BusinessInfo struct {
	Mcc                string     `json:"mcc,omitempty"`
	Websites           []*WebSite `json:"websites,omitempty"`
	EnglishName        string     `json:"englishName,omitempty"`
	DoingBusinessAs    string     `json:"doingBusinessAs,omitempty"`
	MainSalesCountry   string     `json:"mainSalesCountry,omitempty"`
	AppName            string     `json:"appName,omitempty"`
	ServiceDescription string     `json:"serviceDescription,omitempty"`
}
