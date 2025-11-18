package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipaySubmitAttachmentResponse struct {
	response.AlipayResponse
	Result                    *model.Result        `json:"result,omitempty"`
	SubmitAttachmentRequestId string               `json:"submitAttachmentRequestId,omitempty"`
	AttachmentType            model.AttachmentType `json:"attachmentType,omitempty"`
	AttachmentKey             string               `json:"attachmentKey,omitempty"`
}
