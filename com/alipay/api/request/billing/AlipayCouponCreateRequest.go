package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCouponCreateRequest struct {
	CouponRequestId string                       `json:"couponRequestId,omitempty"`
	CouponName      string                       `json:"couponName,omitempty"`
	DiscountType    string                       `json:"discountType,omitempty"`
	PercentOff      string                       `json:"percentOff,omitempty"`
	AmountOff       *model.Amount                `json:"amountOff,omitempty"`
	DurationType    string                       `json:"durationType,omitempty"`
	DurationValue   int32                        `json:"durationValue,omitempty"`
	DurationUnit    string                       `json:"durationUnit,omitempty"`
	MaxRedemptions  int32                        `json:"maxRedemptions,omitempty"`
	RedeemBy        string                       `json:"redeemBy,omitempty"`
	AppliesTo       *model.CouponCreateAppliesTo `json:"appliesTo,omitempty"`
	Metadata        map[string]string            `json:"metadata,omitempty"`
	PromotionCodes  []*model.PromotionCode       `json:"promotionCodes,omitempty"`
}

func NewAlipayCouponCreateRequest() (*request.AlipayRequest, *AlipayCouponCreateRequest) {
	alipayCouponCreateRequest := &AlipayCouponCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCouponCreateRequest, "/ams/api/v1/billing/coupon/create", &responseBilling.AlipayCouponCreateResponse{})
	return alipayRequest, alipayCouponCreateRequest
}

func (alipayCouponCreateRequest *AlipayCouponCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCouponCreateRequest, "/ams/api/v1/billing/coupon/create", &responseBilling.AlipayCouponCreateResponse{})
}
