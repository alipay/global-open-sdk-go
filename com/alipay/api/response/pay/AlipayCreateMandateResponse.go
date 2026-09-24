package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateMandateResponse struct {
	response.AlipayResponse
	Result           *model.Result `json:"result,omitempty"`
	MandateRequestId string        `json:"mandateRequestId,omitempty"`
	MandateId        string        `json:"mandateId,omitempty"`
	ValidFrom        string        `json:"validFrom,omitempty"`
	ValidUntil       string        `json:"validUntil,omitempty"`
}
