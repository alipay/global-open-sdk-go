package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCustomerCreatePortalLinkResponse struct {
	response.AlipayResponse
	Result       *model.Result     `json:"result,omitempty"`
	Token        string            `json:"token,omitempty"`
	PortalUrl    string            `json:"portalUrl,omitempty"`
	ExpiresAt    string            `json:"expiresAt,omitempty"`
	SendStatus   string            `json:"sendStatus,omitempty"`
	Success      bool              `json:"success,omitempty"`
	ErrorContext *model.ErrorStack `json:"errorContext,omitempty"`
}
