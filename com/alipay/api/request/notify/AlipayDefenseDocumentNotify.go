package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayDefenseDocumentNotify struct {
	AlipayNotify
	// Dispute ID.
	DisputeId string `json:"disputeId,omitempty"`
	// Notification time as an ISO 8601 string with timezone.
	NotifyTime string `json:"notifyTime,omitempty"`
	// Metadata of the generated material; the file body is not included.
	DefenseFile *model.DisputeDefenseFile `json:"defenseFile,omitempty"`
	// Optional predicted win-rate level of this material.
	WinRateLevel string `json:"winRateLevel,omitempty"`
	// Optional evidence tags mapped to requirement descriptions.
	RecommendedEvidence map[string]string `json:"recommendedEvidence,omitempty"`
	// Optional suggestions for improving the defense material.
	SuggestionsForImprovement string `json:"suggestionsForImprovement,omitempty"`
}
