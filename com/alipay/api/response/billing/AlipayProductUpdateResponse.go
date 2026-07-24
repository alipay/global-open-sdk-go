package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayProductUpdateResponse struct {
	response.AlipayResponse
	Result        *model.Result     `json:"result,omitempty"`
	ProductId     string            `json:"productId,omitempty"`
	Name          string            `json:"name,omitempty"`
	Type          string            `json:"type,omitempty"`
	Description   string            `json:"description,omitempty"`
	Images        []string          `json:"images,omitempty"`
	UnitLabel     string            `json:"unitLabel,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
	Active        bool              `json:"active,omitempty"`
	CreatedAt     string            `json:"createdAt,omitempty"`
	DeactivatedAt string            `json:"deactivatedAt,omitempty"`
	UpdatedAt     string            `json:"updatedAt,omitempty"`
}
