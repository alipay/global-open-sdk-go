package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceExportResponse struct {
	response.AlipayResponse
	Result     *model.Result `json:"result,omitempty"`
	FileFormat string        `json:"fileFormat,omitempty"`
	ExpiresAt  string        `json:"expiresAt,omitempty"`
	FileUrl    string        `json:"fileUrl,omitempty"`
	FileSize   int64         `json:"fileSize,omitempty"`
	FileName   string        `json:"fileName,omitempty"`
	Mode       string        `json:"mode,omitempty"`
}
