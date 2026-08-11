package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCouponUpdateResponse struct {
	response.AlipayResponse
	Result   *model.ResultInfo `json:"result,omitempty"`
	CouponId string            `json:"couponId,omitempty"`
	Status   string            `json:"status,omitempty"`
}
