package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireCardTransactionLifecycleDetailResponse struct {
	response.AlipayResponse
	Result    *model.Result                         `json:"result,omitempty"`
	Lifecycle *model.CardTransactionLifecycleDetail `json:"lifecycle,omitempty"`
}
