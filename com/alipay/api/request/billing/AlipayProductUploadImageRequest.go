package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayProductUploadImageRequest struct {
	ProductId string `json:"productId,omitempty"`
	ImageFile string `json:"imageFile,omitempty"`
}

func NewAlipayProductUploadImageRequest() (*request.AlipayRequest, *AlipayProductUploadImageRequest) {
	alipayProductUploadImageRequest := &AlipayProductUploadImageRequest{}
	alipayRequest := request.NewAlipayRequest(alipayProductUploadImageRequest, "/ams/api/v1/billing/product/uploadImage", &responseBilling.AlipayProductUploadImageResponse{})
	return alipayRequest, alipayProductUploadImageRequest
}

func (alipayProductUploadImageRequest *AlipayProductUploadImageRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayProductUploadImageRequest, "/ams/api/v1/billing/product/uploadImage", &responseBilling.AlipayProductUploadImageResponse{})
}
