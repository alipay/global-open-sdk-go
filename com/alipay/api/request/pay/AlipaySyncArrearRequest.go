package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipaySyncArrearRequest struct {
	PaymentId        string `json:"paymentId,omitempty"`
	PaymentRequestId string `json:"paymentRequestId,omitempty"`
}

func (alipayCaptureRequest *AlipaySyncArrearRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCaptureRequest, model.SYNC_ARREAR_PATH, &responsePay.AlipaySyncArrearResponse{})
}

func NewAlipaySyncArrearRequest() (*request.AlipayRequest, *AlipaySyncArrearRequest) {
	alipaySyncArrearRequest := &AlipaySyncArrearRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySyncArrearRequest, model.SYNC_ARREAR_PATH, &responsePay.AlipaySyncArrearResponse{})
	return alipayRequest, alipaySyncArrearRequest
}
