package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquiryRateResponse struct {
	response.AlipayResponse
	RateResultList []*model.RateResult `json:"rateResultList,omitempty"`
	Result         *model.Result       `json:"result,omitempty"`
}
