package responseCustoms

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCustomsQueryResponse struct {
	response.AlipayResponse
	DeclarationRequestsNotFound []string                  `json:"declarationRequestsNotFound"`
	DeclarationRecords          []model.DeclarationRecord `json:"declarationRecords"`
}
