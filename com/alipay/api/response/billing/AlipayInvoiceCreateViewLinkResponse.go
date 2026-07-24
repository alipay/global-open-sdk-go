package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceCreateViewLinkResponse struct {
	response.AlipayResponse
	Result    *model.InvoiceCreateViewLinkResult `json:"result,omitempty"`
	Token     string                             `json:"token,omitempty"`
	ViewUrl   string                             `json:"viewUrl,omitempty"`
	ExpiresAt string                             `json:"expiresAt,omitempty"`
}
