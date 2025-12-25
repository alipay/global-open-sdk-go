package model

type AbaCard struct {
	AssetId      string `json:"assetId,omitempty"`
	CardNickName string `json:"cardNickName,omitempty"`
	MaskedCardNo string `json:"maskedCardNo,omitempty"`
	CardStatus   string `json:"cardStatus,omitempty"`
	CardBrand    string `json:"cardBrand,omitempty"`
	CreatedTime  string `json:"createdTime,omitempty"`
	UpdatedTime  string `json:"updatedTime,omitempty"`
}
