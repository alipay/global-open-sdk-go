package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayProductUploadImageResponse struct {
	response.AlipayResponse
	Result    *model.Result `json:"result,omitempty"`
	ImageUrl  string        `json:"imageUrl,omitempty"`
	ImageName string        `json:"imageName,omitempty"`
}
