package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayQueryTradeBillListRequest struct {
	MerchantId     string `json:"merchantId,omitempty"`
	CurrentPage    int32  `json:"currentPage,omitempty"`
	PageSize       int32  `json:"pageSize,omitempty"`
	StartTimestamp int64  `json:"startTimestamp,omitempty"`
	EndTimestamp   int64  `json:"endTimestamp,omitempty"`
	TimeZone       string `json:"timeZone,omitempty"`
}

func NewAlipayQueryTradeBillListRequest() (*request.AlipayRequest, *AlipayQueryTradeBillListRequest) {
	alipayQueryTradeBillListRequest := &AlipayQueryTradeBillListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayQueryTradeBillListRequest, "/ams/api/v1/adaptor/queryTradeBillList", &responsePay.AlipayQueryTradeBillListResponse{})
	return alipayRequest, alipayQueryTradeBillListRequest
}

func (alipayQueryTradeBillListRequest *AlipayQueryTradeBillListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayQueryTradeBillListRequest, "/ams/api/v1/adaptor/queryTradeBillList", &responsePay.AlipayQueryTradeBillListResponse{})
}
