package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireCardSensitiveInfoResponse struct {
	response.AlipayResponse
	AuthorizationControl *model.AlipayInquireCardSensitiveInfoResponseAuthorizationControl `json:"authorizationControl,omitempty"`
	Result               *model.Result                                                     `json:"result,omitempty"`
	AssetId              string                                                            `json:"assetId,omitempty"`
	Cvv                  string                                                            `json:"cvv,omitempty"`
	CardNo               string                                                            `json:"cardNo,omitempty"`
	ExpiredMonth         string                                                            `json:"expiredMonth,omitempty"`
	ExpiredYear          string                                                            `json:"expiredYear,omitempty"`
}
