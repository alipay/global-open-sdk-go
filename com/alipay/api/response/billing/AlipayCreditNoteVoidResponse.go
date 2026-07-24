package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreditNoteVoidResponse struct {
	response.AlipayResponse
	Result       *model.Result `json:"result,omitempty"`
	CreditNoteId string        `json:"creditNoteId,omitempty"`
	Status       string        `json:"status,omitempty"`
	VoidedAt     string        `json:"voidedAt,omitempty"`
}
