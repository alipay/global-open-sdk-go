package risk

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseRisk "github.com/alipay/global-open-sdk-go/com/alipay/api/response/risk"
)

type RiskReportRequest struct {
	ReferenceTransactionId string `json:"referenceTransactionId,omitempty"`
	ReportReason           string `json:"reportReason,omitempty"`
	RiskType               string `json:"riskType,omitempty"`
	RiskOccurrenceTime     string `json:"riskOccurrenceTime,omitempty"`
}

func NewRiskReportRequest() (*request.AlipayRequest, *RiskReportRequest) {
	riskReportRequest := &RiskReportRequest{}
	alipayRequest := request.NewAlipayRequest(riskReportRequest, model.RISK_REPORT_PATH, &responseRisk.RiskReportResponse{})
	return alipayRequest, riskReportRequest
}
