package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayDefenseStrategySuggestionNotify struct {
	AlipayNotify
	// Dispute ID.
	DisputeId string `json:"disputeId,omitempty"`
	// Notification time as an ISO 8601 string with timezone.
	NotifyTime string `json:"notifyTime,omitempty"`
	// Predicted win-rate level: HIGH, MEDIUM or LOW.
	WinRateLevel string `json:"winRateLevel,omitempty"`
	// Recommended evidence items, from 1 to 32 entries.
	EvidenceList []model.DisputeEvidenceRecommendation `json:"evidenceList,omitempty"`
	// Optional cardholder dispute reason.
	CardholderDisputeReason string `json:"cardholderDisputeReason,omitempty"`
	// Suggested defense strategy.
	DefenseStrategyDescription string `json:"defenseStrategyDescription,omitempty"`
}
