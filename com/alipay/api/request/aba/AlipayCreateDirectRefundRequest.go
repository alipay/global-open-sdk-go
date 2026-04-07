package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayCreateDirectRefundRequest struct {
	PaymentId         string                  `json:"paymentId,omitempty"`
	RefundRequestId   string                  `json:"refundRequestId,omitempty"`
	ReferenceRefundId string                  `json:"referenceRefundId,omitempty"`
	RefundFromMethod  *model.RefundFromMethod `json:"refundFromMethod,omitempty"`
	RefundFromAmount  *model.Amount           `json:"refundFromAmount,omitempty"`
	Memo              string                  `json:"memo,omitempty"`
	Remark            string                  `json:"remark,omitempty"`
	RefundReason      string                  `json:"refundReason,omitempty"`
	RefundNotifyUrl   string                  `json:"refundNotifyUrl,omitempty"`
}

func NewAlipayCreateDirectRefundRequest() (*request.AlipayRequest, *AlipayCreateDirectRefundRequest) {
	alipayCreateDirectRefundRequest := &AlipayCreateDirectRefundRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateDirectRefundRequest, "null", &responseAba.AlipayCreateDirectRefundResponse{})
	return alipayRequest, alipayCreateDirectRefundRequest
}

func (alipayCreateDirectRefundRequest *AlipayCreateDirectRefundRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateDirectRefundRequest, "null", &responseAba.AlipayCreateDirectRefundResponse{})
}
