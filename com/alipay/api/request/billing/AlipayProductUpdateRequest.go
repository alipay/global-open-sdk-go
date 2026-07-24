package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayProductUpdateRequest struct {
	ProductId            string            `json:"productId,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Description          string            `json:"description,omitempty"`
	Images               []string          `json:"images,omitempty"`
	UnitLabel            string            `json:"unitLabel,omitempty"`
	Metadata             map[string]string `json:"metadata,omitempty"`
	MetadataKeysToRemove []string          `json:"metadataKeysToRemove,omitempty"`
	Active               bool              `json:"active,omitempty"`
}

func NewAlipayProductUpdateRequest() (*request.AlipayRequest, *AlipayProductUpdateRequest) {
	alipayProductUpdateRequest := &AlipayProductUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayProductUpdateRequest, "/ams/api/v1/billing/product/update", &responseBilling.AlipayProductUpdateResponse{})
	return alipayRequest, alipayProductUpdateRequest
}

func (alipayProductUpdateRequest *AlipayProductUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayProductUpdateRequest, "/ams/api/v1/billing/product/update", &responseBilling.AlipayProductUpdateResponse{})
}
