package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireCardResponse struct {
	response.AlipayResponse
	Result            *model.Result    `json:"result,omitempty"`
	TotalCount        int32            `json:"totalCount,omitempty"`
	TotalPageNumber   int32            `json:"totalPageNumber,omitempty"`
	CurrentPageNumber int32            `json:"currentPageNumber,omitempty"`
	CardList          []*model.AbaCard `json:"cardList,omitempty"`
}
