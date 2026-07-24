package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPriceUpdateResponse struct {
	response.AlipayResponse
	Result           *model.Result            `json:"result,omitempty"`
	PriceId          string                   `json:"priceId,omitempty"`
	ProductId        string                   `json:"productId,omitempty"`
	Name             string                   `json:"name,omitempty"`
	PricingModel     string                   `json:"pricingModel,omitempty"`
	UsageType        string                   `json:"usageType,omitempty"`
	UnitLabel        string                   `json:"unitLabel,omitempty"`
	MeterId          string                   `json:"meterId,omitempty"`
	UnitAmount       *model.Amount            `json:"unitAmount,omitempty"`
	Recurring        *model.RecurringSettings `json:"recurring,omitempty"`
	Active           bool                     `json:"active,omitempty"`
	IncludedQuantity int32                    `json:"includedQuantity,omitempty"`
	TiersMode        string                   `json:"tiersMode,omitempty"`
	Tiers            []*model.Tier            `json:"tiers,omitempty"`
	Metadata         map[string]string        `json:"metadata,omitempty"`
	CreatedAt        string                   `json:"createdAt,omitempty"`
	DeactivatedAt    string                   `json:"deactivatedAt,omitempty"`
	UpdatedAt        string                   `json:"updatedAt,omitempty"`
}
