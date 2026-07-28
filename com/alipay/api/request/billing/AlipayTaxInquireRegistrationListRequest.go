package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxInquireRegistrationListRequest struct {
	Status      string `json:"status,omitempty"`
	CurrentPage int32  `json:"currentPage,omitempty"`
	PageSize    int32  `json:"pageSize,omitempty"`
}

func NewAlipayTaxInquireRegistrationListRequest() (*request.AlipayRequest, *AlipayTaxInquireRegistrationListRequest) {
	alipayTaxInquireRegistrationListRequest := &AlipayTaxInquireRegistrationListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxInquireRegistrationListRequest, "/ams/api/v1/tax/inquireRegistrationList", &responseBilling.AlipayTaxInquireRegistrationListResponse{})
	return alipayRequest, alipayTaxInquireRegistrationListRequest
}

func (alipayTaxInquireRegistrationListRequest *AlipayTaxInquireRegistrationListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxInquireRegistrationListRequest, "/ams/api/v1/tax/inquireRegistrationList", &responseBilling.AlipayTaxInquireRegistrationListResponse{})
}
