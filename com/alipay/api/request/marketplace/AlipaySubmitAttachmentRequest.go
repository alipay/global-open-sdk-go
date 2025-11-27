package marketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseMarketplace "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace"
)

type AlipaySubmitAttachmentRequest struct {
	SubmitAttachmentRequestId string               `json:"submitAttachmentRequestId,omitempty"`
	AttachmentType            model.AttachmentType `json:"attachmentType,omitempty"`
	FileSha256                string               `json:"fileSha256,omitempty"`
}

func NewAlipaySubmitAttachmentRequest() (*request.AlipayRequest, *AlipaySubmitAttachmentRequest) {
	alipaySubmitAttachmentRequest := &AlipaySubmitAttachmentRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySubmitAttachmentRequest, "/ams/api/open/openapiv2_file/v1/business/attachment/submitAttachment", &responseMarketplace.AlipaySubmitAttachmentResponse{})
	return alipayRequest, alipaySubmitAttachmentRequest
}

func (alipaySubmitAttachmentRequest *AlipaySubmitAttachmentRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipaySubmitAttachmentRequest, "/ams/api/open/openapiv2_file/v1/business/attachment/submitAttachment", &responseMarketplace.AlipaySubmitAttachmentResponse{})
}
