package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayPayQueryRequest struct {
	PaymentRequestId  string `json:"paymentRequestId,omitempty"`
	PaymentId         string `json:"paymentId,omitempty"`
	MerchantAccountId string `json:"MerchantAccountId,omitempty"`
}

func (alipayPayQueryRequest *AlipayPayQueryRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPayQueryRequest, model.INQUIRY_PAYMENT_PATH, &responsePay.AlipayPayQueryResponse{})
}

func NewAlipayPayQueryRequest() (*request.AlipayRequest, *AlipayPayQueryRequest) {
	alipayPayQueryRequest := &AlipayPayQueryRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPayQueryRequest, model.INQUIRY_PAYMENT_PATH, &responsePay.AlipayPayQueryResponse{})
	return alipayRequest, alipayPayQueryRequest
}
