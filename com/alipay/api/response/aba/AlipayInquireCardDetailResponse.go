package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireCardDetailResponse struct {
	response.AlipayResponse
	Result               *model.Result               `json:"result,omitempty"`
	AssetId              string                      `json:"assetId,omitempty"`
	CardNickName         string                      `json:"cardNickName,omitempty"`
	CardStatus           string                      `json:"cardStatus,omitempty"`
	MaskedCardNo         string                      `json:"maskedCardNo,omitempty"`
	CardBrand            string                      `json:"cardBrand,omitempty"`
	CreatedTime          string                      `json:"createdTime,omitempty"`
	UpdatedTime          string                      `json:"updatedTime,omitempty"`
	Purpose              string                      `json:"purpose,omitempty"`
	Note                 string                      `json:"note,omitempty"`
	Metadata             map[string]string           `json:"metadata,omitempty"`
	AuthorizationControl *model.AuthorizationControl `json:"authorizationControl,omitempty"`
	Cardholderinfo       *model.CardholderInfo       `json:"cardholderinfo,omitempty"`
}
