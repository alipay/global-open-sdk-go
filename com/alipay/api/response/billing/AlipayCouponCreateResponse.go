package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCouponCreateResponse struct {
	response.AlipayResponse
	Result         *model.ResultInfo            `json:"result,omitempty"`
	CouponId       string                       `json:"couponId,omitempty"`
	CouponName     string                       `json:"couponName,omitempty"`
	DiscountType   string                       `json:"discountType,omitempty"`
	Status         string                       `json:"status,omitempty"`
	PromotionCodes []*model.CouponPromotionCode `json:"promotionCodes,omitempty"`
}
