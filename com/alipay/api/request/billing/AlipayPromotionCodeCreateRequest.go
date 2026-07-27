package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayPromotionCodeCreateRequest struct {
	PromotionCodeRequestId string                              `json:"promotionCodeRequestId,omitempty"`
	CouponId               string                              `json:"couponId,omitempty"`
	Code                   string                              `json:"code,omitempty"`
	MaxRedeemSize          int32                               `json:"maxRedeemSize,omitempty"`
	ExpiryTime             string                              `json:"expiryTime,omitempty"`
	MinAmount              *model.PromotionCodeCreateMinAmount `json:"minAmount,omitempty"`
	OneTimeOnly            bool                                `json:"oneTimeOnly,omitempty"`
	CustomerId             string                              `json:"customerId,omitempty"`
	Metadata               map[string]string                   `json:"metadata,omitempty"`
}

func NewAlipayPromotionCodeCreateRequest() (*request.AlipayRequest, *AlipayPromotionCodeCreateRequest) {
	alipayPromotionCodeCreateRequest := &AlipayPromotionCodeCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPromotionCodeCreateRequest, "/ams/api/v1/billing/promotionCode/create", &responseBilling.AlipayPromotionCodeCreateResponse{})
	return alipayRequest, alipayPromotionCodeCreateRequest
}

func (alipayPromotionCodeCreateRequest *AlipayPromotionCodeCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPromotionCodeCreateRequest, "/ams/api/v1/billing/promotionCode/create", &responseBilling.AlipayPromotionCodeCreateResponse{})
}
