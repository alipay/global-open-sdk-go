package risk

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseRisk "github.com/alipay/global-open-sdk-go/com/alipay/api/response/risk"
)

type SendRefundResultRequest struct {
	ReferenceTransactionId string                `json:"referenceTransactionId,omitempty"`
	ReferenceRefundId      string                `json:"referenceRefundId,omitempty"`
	ActualRefundAmount     *model.Amount         `json:"actualRefundAmount,omitempty"`
	RefundRecords          []*model.RefundRecord `json:"refundRecords,omitempty"`
}

func NewSendRefundResultRequest() (*request.AlipayRequest, *SendRefundResultRequest) {
	sendRefundResultRequest := &SendRefundResultRequest{}
	alipayRequest := request.NewAlipayRequest(sendRefundResultRequest, model.RISK_SEND_REFUND_RESULT_PATH, &responseRisk.SendRefundResultResponse{})
	return alipayRequest, sendRefundResultRequest
}
