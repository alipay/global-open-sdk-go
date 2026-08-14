package model

type CardTransactionLifecycle struct {
	LifecycleId           string  `json:"lifecycleId,omitempty"`
	LatestEventType       string  `json:"latestEventType,omitempty"`
	LatestEventStatus     string  `json:"latestEventStatus,omitempty"`
	LastUpdateTime        string  `json:"lastUpdateTime,omitempty"`
	TransactionTime       string  `json:"transactionTime,omitempty"`
	TotalBillingAmount    *Amount `json:"totalBillingAmount,omitempty"`
	TotalAuthAmount       *Amount `json:"totalAuthAmount,omitempty"`
	TotalCancelAmount     *Amount `json:"totalCancelAmount,omitempty"`
	TotalRefundAmount     *Amount `json:"totalRefundAmount,omitempty"`
	TotalChargebackAmount *Amount `json:"totalChargebackAmount,omitempty"`
	AssetId               string  `json:"assetId,omitempty"`
	MaskedCardNo          string  `json:"maskedCardNo,omitempty"`
}
