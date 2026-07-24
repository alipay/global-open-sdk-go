package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayPriceCreateRequest struct {
	PriceRequestId   string                   `json:"priceRequestId,omitempty"`
	ProductId        string                   `json:"productId,omitempty"`
	Name             string                   `json:"name,omitempty"`
	PricingModel     string                   `json:"pricingModel,omitempty"`
	UsageType        string                   `json:"usageType,omitempty"`
	UnitAmount       *model.Amount            `json:"unitAmount,omitempty"`
	UnitLabel        string                   `json:"unitLabel,omitempty"`
	MeterId          string                   `json:"meterId,omitempty"`
	Recurring        *model.RecurringSettings `json:"recurring,omitempty"`
	IncludedQuantity int32                    `json:"includedQuantity,omitempty"`
	TiersMode        string                   `json:"tiersMode,omitempty"`
	Tiers            []*model.Tier            `json:"tiers,omitempty"`
	Metadata         map[string]string        `json:"metadata,omitempty"`
}

func NewAlipayPriceCreateRequest() (*request.AlipayRequest, *AlipayPriceCreateRequest) {
	alipayPriceCreateRequest := &AlipayPriceCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPriceCreateRequest, "/ams/api/v1/billing/price/create", &responseBilling.AlipayPriceCreateResponse{})
	return alipayRequest, alipayPriceCreateRequest
}

func (alipayPriceCreateRequest *AlipayPriceCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPriceCreateRequest, "/ams/api/v1/billing/price/create", &responseBilling.AlipayPriceCreateResponse{})
}
