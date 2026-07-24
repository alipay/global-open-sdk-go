package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPromotionCodeInquireDetailsResponse struct {
	response.AlipayResponse
	Result                 *model.Result                               `json:"result,omitempty"`
	PromotionCodeId        string                                      `json:"promotionCodeId,omitempty"`
	PromotionCodeRequestId string                                      `json:"promotionCodeRequestId,omitempty"`
	Code                   string                                      `json:"code,omitempty"`
	CouponId               string                                      `json:"couponId,omitempty"`
	Status                 string                                      `json:"status,omitempty"`
	MaxRedemptions         int32                                       `json:"maxRedemptions,omitempty"`
	RedeemedCount          int32                                       `json:"redeemedCount,omitempty"`
	ExpiryTime             string                                      `json:"expiryTime,omitempty"`
	MinAmount              *model.PromotionCodeInquireDetailsMinAmount `json:"minAmount,omitempty"`
	OneTimeOnly            bool                                        `json:"oneTimeOnly,omitempty"`
	CustomerId             string                                      `json:"customerId,omitempty"`
	Metadata               map[string]string                           `json:"metadata,omitempty"`
	CreateTime             string                                      `json:"createTime,omitempty"`
	UpdateTime             string                                      `json:"updateTime,omitempty"`
}
