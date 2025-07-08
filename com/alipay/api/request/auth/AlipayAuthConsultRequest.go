package auth

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAuth "github.com/alipay/global-open-sdk-go/com/alipay/api/response/auth"
)

type AlipayAuthConsultRequest struct {
	CustomerBelongsTo model.CustomerBelongsTo `json:"customerBelongsTo,omitempty"`
	AuthClientId      string                  `json:"authClientId,omitempty"`
	AuthRedirectUrl   string                  `json:"authRedirectUrl,omitempty"`
	Scopes            []model.ScopeType       `json:"scopes,omitempty"`
	AuthState         string                  `json:"authState,omitempty"`
	TerminalType      model.TerminalType      `json:"terminalType,omitempty"`
	OsType            model.OsType            `json:"osType,omitempty"`
	OsVersion         string                  `json:"osVersion,omitempty"`
	ExtendInfo        string                  `json:"extendInfo,omitempty"`
	MerchantRegion    string                  `json:"merchantRegion,omitempty"`
	RecurringPayment  bool                    `json:"recurringPayment,omitempty"`
	AuthMetaData      *model.AuthMetaData     `json:"authMetaData,omitempty"`
}

func NewAlipayAuthConsultRequest() (*request.AlipayRequest, *AlipayAuthConsultRequest) {
	alipayAuthConsultRequest := &AlipayAuthConsultRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAuthConsultRequest, "/ams/api/v1/authorizations/consult", &responseAuth.AlipayAuthConsultResponse{})
	return alipayRequest, alipayAuthConsultRequest
}

func (alipayAuthConsultRequest *AlipayAuthConsultRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayAuthConsultRequest, "/ams/api/v1/authorizations/consult", &responseAuth.AlipayAuthConsultResponse{})
}
