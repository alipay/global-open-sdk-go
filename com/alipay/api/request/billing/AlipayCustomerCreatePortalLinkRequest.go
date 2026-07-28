package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCustomerCreatePortalLinkRequest struct {
	CustomerId string   `json:"customerId,omitempty"`
	Email      string   `json:"email,omitempty"`
	ExpiryDays int32    `json:"expiryDays,omitempty"`
	Features   []string `json:"features,omitempty"`
	AutoSend   bool     `json:"autoSend,omitempty"`
	SettingId  string   `json:"settingId,omitempty"`
}

func NewAlipayCustomerCreatePortalLinkRequest() (*request.AlipayRequest, *AlipayCustomerCreatePortalLinkRequest) {
	alipayCustomerCreatePortalLinkRequest := &AlipayCustomerCreatePortalLinkRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCustomerCreatePortalLinkRequest, "/ams/api/v1/billing/customer/createPortalLink", &responseBilling.AlipayCustomerCreatePortalLinkResponse{})
	return alipayRequest, alipayCustomerCreatePortalLinkRequest
}

func (alipayCustomerCreatePortalLinkRequest *AlipayCustomerCreatePortalLinkRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCustomerCreatePortalLinkRequest, "/ams/api/v1/billing/customer/createPortalLink", &responseBilling.AlipayCustomerCreatePortalLinkResponse{})
}
