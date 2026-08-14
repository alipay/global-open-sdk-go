package defaultAlipayClient

import (
	"strings"

	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
)

type requestTransportRoute struct {
	method string
	path   string
}

var sessionHTTP2Routes = []requestTransportRoute{
	{method: model.MethodPost, path: "/ams/api/v1/meter/uploadEvent"},
}

func requiresSessionHTTP2(alipayRequest *request.AlipayRequest) bool {
	if alipayRequest == nil {
		return false
	}
	for _, route := range sessionHTTP2Routes {
		if strings.EqualFold(alipayRequest.HttpMethod, route.method) && alipayRequest.Path == route.path {
			return true
		}
	}
	return false
}
