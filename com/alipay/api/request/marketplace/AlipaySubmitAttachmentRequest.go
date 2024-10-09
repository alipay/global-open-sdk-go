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
	alipayRequest := request.NewAlipayRequest(alipaySubmitAttachmentRequest, model.MARKETPLACE_SUBMITATTACHMENT_PATH, &responseMarketplace.AlipaySubmitAttachmentResponse{})
	return alipayRequest, alipaySubmitAttachmentRequest
}
