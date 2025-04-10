package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayUploadInvoiceShippingFileRequest struct {
	PaymentRequestId string `json:"paymentRequestId,omitempty"`
	MerchantId       string `json:"merchantId,omitempty"`
	FileId           string `json:"fileId,omitempty"`
	UploadFile       string `json:"uploadFile,omitempty"`
	FileType         string `json:"fileType,omitempty"`
	FileName         string `json:"fileName,omitempty"`
}

func (alipayUploadInvoiceShippingFileRequest *AlipayUploadInvoiceShippingFileRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayUploadInvoiceShippingFileRequest, "/ams/api/v1/payments/uploadInvoiceShippingFile", &responsePay.AlipayUploadInvoiceShippingFileResponse{})
}

func NewAlipayUploadInvoiceShippingFileRequest() (*request.AlipayRequest, *AlipayUploadInvoiceShippingFileRequest) {
	alipayUploadInvoiceShippingFileRequest := &AlipayUploadInvoiceShippingFileRequest{}
	alipayRequest := request.NewAlipayRequest(alipayUploadInvoiceShippingFileRequest, "/ams/api/v1/payments/uploadInvoiceShippingFile", &responsePay.AlipayUploadInvoiceShippingFileResponse{})
	return alipayRequest, alipayUploadInvoiceShippingFileRequest
}
