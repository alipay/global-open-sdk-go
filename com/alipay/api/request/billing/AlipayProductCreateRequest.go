package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayProductCreateRequest struct {
	ProductRequestId string            `json:"productRequestId,omitempty"`
	Name             string            `json:"name,omitempty"`
	Type             string            `json:"type,omitempty"`
	Description      string            `json:"description,omitempty"`
	Images           []string          `json:"images,omitempty"`
	UnitLabel        string            `json:"unitLabel,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
}

func NewAlipayProductCreateRequest() (*request.AlipayRequest, *AlipayProductCreateRequest) {
	alipayProductCreateRequest := &AlipayProductCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayProductCreateRequest, "/ams/api/v1/billing/product/create", &responseBilling.AlipayProductCreateResponse{})
	return alipayRequest, alipayProductCreateRequest
}

func (alipayProductCreateRequest *AlipayProductCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayProductCreateRequest, "/ams/api/v1/billing/product/create", &responseBilling.AlipayProductCreateResponse{})
}
