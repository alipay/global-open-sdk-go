package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayRefundRequest struct {
	RefundRequestId       string               `json:"refundRequestId,omitempty"`
	PaymentId             string               `json:"paymentId,omitempty"`
	ReferenceRefundId     string               `json:"referenceRefundId,omitempty"`
	RefundAmount          *model.Amount        `json:"refundAmount,omitempty"`
	RefundReason          string               `json:"refundReason,omitempty"`
	RefundNotifyUrl       string               `json:"refundNotifyUrl,omitempty"`
	IsAsyncRefund         *bool                `json:"isAsyncRefund,omitempty"`
	ExtendInfo            string               `json:"extendInfo,omitempty"`
	RefundDetails         []model.RefundDetail `json:"refundDetails,omitempty"`
	RefundSourceAccountNo string               `json:"refundSourceAccountNo,omitempty"`
}

func (alipayRefundRequest *AlipayRefundRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayRefundRequest, model.REFUND_PATH, &responsePay.AlipayRefundResponse{})
}

func NewAlipayRefundRequest() (*request.AlipayRequest, *AlipayRefundRequest) {
	alipayRefundRequest := &AlipayRefundRequest{}
	alipayRequest := request.NewAlipayRequest(alipayRefundRequest, model.REFUND_PATH, &responsePay.AlipayRefundResponse{})
	return alipayRequest, alipayRefundRequest
}
