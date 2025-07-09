package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayRetrievePaymentSessionResponse struct {
	response.AlipayResponse
	Order            model.Order             `json:"order"`
	PromotionResults []model.PromotionResult `json:"promotionResults"`
	CustomizedInfo   string                  `json:"customizedInfo"`
}
