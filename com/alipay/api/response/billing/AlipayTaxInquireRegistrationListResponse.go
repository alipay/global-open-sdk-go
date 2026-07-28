package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayTaxInquireRegistrationListResponse struct {
	response.AlipayResponse
	Result        *model.Result            `json:"result,omitempty"`
	Registrations []*model.TaxRegistration `json:"registrations,omitempty"`
	Paginator     *model.Paginator         `json:"paginator,omitempty"`
}
