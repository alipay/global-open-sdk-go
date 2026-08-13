package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayMeterInquireListResponse struct {
	response.AlipayResponse
	Result     *model.Result  `json:"result,omitempty"`
	PageNum    int32          `json:"pageNum,omitempty"`
	PageSize   int32          `json:"pageSize,omitempty"`
	TotalCount int32          `json:"totalCount,omitempty"`
	Meters     []*model.Meter `json:"meters,omitempty"`
}
