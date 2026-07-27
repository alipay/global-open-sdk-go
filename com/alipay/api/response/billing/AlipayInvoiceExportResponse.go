package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceExportResponse struct {
	response.AlipayResponse
	Result      *model.Result `json:"result,omitempty"`
	Status      string        `json:"status,omitempty"`
	DownloadUrl string        `json:"downloadUrl,omitempty"`
	ExpiresAt   string        `json:"expiresAt,omitempty"`
}
