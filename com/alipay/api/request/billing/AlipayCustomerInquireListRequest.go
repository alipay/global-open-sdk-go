package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCustomerInquireListRequest struct {
	StartingAfter  string   `json:"startingAfter,omitempty"`
	EndingBefore   string   `json:"endingBefore,omitempty"`
	Limit          int32    `json:"limit,omitempty"`
	IncludeTotal   bool     `json:"includeTotal,omitempty"`
	Status         string   `json:"status,omitempty"`
	Email          string   `json:"email,omitempty"`
	PhoneNo        string   `json:"phoneNo,omitempty"`
	CountryCode    string   `json:"countryCode,omitempty"`
	GmtCreateEnd   string   `json:"gmtCreateEnd,omitempty"`
	CustomerId     string   `json:"customerId,omitempty"`
	Country        []string `json:"country,omitempty"`
	EmailPrefix    string   `json:"emailPrefix,omitempty"`
	GmtCreateStart string   `json:"gmtCreateStart,omitempty"`
}

func NewAlipayCustomerInquireListRequest() (*request.AlipayRequest, *AlipayCustomerInquireListRequest) {
	alipayCustomerInquireListRequest := &AlipayCustomerInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCustomerInquireListRequest, "/ams/api/v1/billing/customer/inquireList", &responseBilling.AlipayCustomerInquireListResponse{})
	return alipayRequest, alipayCustomerInquireListRequest
}

func (alipayCustomerInquireListRequest *AlipayCustomerInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCustomerInquireListRequest, "/ams/api/v1/billing/customer/inquireList", &responseBilling.AlipayCustomerInquireListResponse{})
}
