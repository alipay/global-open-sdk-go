package model

type CardTransactionEvent struct {
	EventId              string             `json:"eventId,omitempty"`
	LifecycleId          string             `json:"lifecycleId,omitempty"`
	EventType            string             `json:"eventType,omitempty"`
	AuthExpireTime       string             `json:"authExpireTime,omitempty"`
	AuthType             string             `json:"authType,omitempty"`
	AuthCode             string             `json:"authCode,omitempty"`
	FailureReason        string             `json:"failureReason,omitempty"`
	Status               string             `json:"status,omitempty"`
	BalanceType          string             `json:"balanceType,omitempty"`
	TransactionTime      string             `json:"transactionTime,omitempty"`
	BillType             string             `json:"billType,omitempty"`
	OutAmount            *Amount            `json:"outAmount,omitempty"`
	InAmount             *Amount            `json:"inAmount,omitempty"`
	ExchangeCurrencyPair string             `json:"exchangeCurrencyPair,omitempty"`
	ExchangeRate         string             `json:"exchangeRate,omitempty"`
	TransactionAmount    *Amount            `json:"transactionAmount,omitempty"`
	AssetId              string             `json:"assetId,omitempty"`
	MaskedCardNo         string             `json:"maskedCardNo,omitempty"`
	MerchantInfo         *EventMerchantInfo `json:"merchantInfo,omitempty"`
	Metadata             map[string]string  `json:"metadata,omitempty"`
}
