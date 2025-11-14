package model

type WebSite struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
	Desc string `json:"desc,omitempty"`
	Type string `json:"type,omitempty"`
}
