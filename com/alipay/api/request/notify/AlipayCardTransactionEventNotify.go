package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

// AlipayCardTransactionEventNotify is one flat ABA VCC card transaction event callback.
// It follows notification profile v2.1 (document v2.13). Verify the raw-body signature before
// parsing; the merchant handles durable acceptance, eventId deduplication and ACK.
// Amount values remain quoted minor-unit strings of up to 20 digits.
type AlipayCardTransactionEventNotify struct {
	// Stable event ID and the only callback idempotency key. Required; maximum 64 characters.
	EventId string `json:"eventId"`
	// Spend-group ID shared with inquiry. Required; maximum 64 characters.
	LifecycleId string `json:"lifecycleId"`
	// Required event type: AUTH, AUTH_CANCEL, CAPTURE, REFUND, CHARGEBACK or REPAYMENT. Preserve unknown future values.
	EventType string `json:"eventType"`
	// Optional ISO 8601 time of the channel-correlated authorization, maximum 32 characters. Omitted for AUTH or when unavailable; never inferred from transactionTime or authExpireTime.
	AuthTime string `json:"authTime,omitempty"`
	// Optional ISO 8601 authorization expiry time, maximum 32 characters. Sent only for AUTH when available.
	AuthExpireTime string `json:"authExpireTime,omitempty"`
	// Optional AUTH or PRE_AUTH, maximum 32 characters. Sent only for AUTH when available.
	AuthType string `json:"authType,omitempty"`
	// Optional network approval code, maximum 64 characters. Sent only for AUTH when available.
	AuthCode string `json:"authCode,omitempty"`
	// Optional Antom-normalized failure code, maximum 64 characters. Sent only for FAIL when a mapped code exists.
	FailureReason string `json:"failureReason,omitempty"`
	// Required terminal status: SUCCESS or FAIL. CAPTURE is SUCCESS-only. The accepted status never changes.
	Status string `json:"status"`
	// ANTOM_BIZ_ACCOUNT or ANTOM_FUND_ACCOUNT, maximum 32 characters. Required with a directional amount; CAPTURE sends ANTOM_BIZ_ACCOUNT.
	BalanceType string `json:"balanceType,omitempty"`
	// Required business-event time in ISO 8601 format with timezone; maximum 32 characters.
	TransactionTime string `json:"transactionTime"`
	// Required CARD_PAYMENT for AUTH, CAPTURE and REPAYMENT, or CARD_REFUND for AUTH_CANCEL, REFUND and CHARGEBACK.
	BillType string `json:"billType"`
	// Required for successful AUTH, CAPTURE and REPAYMENT. CAPTURE carries value "0" and billing currency, not another debit; omitted otherwise.
	OutAmount *model.Amount `json:"outAmount,omitempty"`
	// Optional billing-currency inflow for successful AUTH_CANCEL, REFUND or CHARGEBACK when available. Omitted for CAPTURE.
	InAmount *model.Amount `json:"inAmount,omitempty"`
	// Optional BASE/QUOTE currency pair, maximum 8 characters. Sent with exchangeRate for a cross-currency transaction when both are available.
	ExchangeCurrencyPair string `json:"exchangeCurrencyPair,omitempty"`
	// Optional positive quoted decimal string, maximum 32 characters. One major BASE unit equals this many QUOTE units; paired with exchangeCurrencyPair.
	ExchangeRate string `json:"exchangeRate,omitempty"`
	// Required current-event amount in original transaction currency. For CAPTURE this is the captured amount, not the zero additional outflow.
	TransactionAmount *model.Amount `json:"transactionAmount"`
	// Required merchant-owned card asset ID; maximum 32 characters.
	AssetId string `json:"assetId"`
	// Optional PCI-safe card mask exposing at most the last four digits; maximum 32 characters.
	MaskedCardNo string `json:"maskedCardNo,omitempty"`
	// Optional original-transaction program code matching ^[A-Z0-9]{3}$. Extensible string, not a fixed enum or rebate rate; later card-profile changes do not rewrite it.
	CardProduct string `json:"cardProduct,omitempty"`
	// Optional approved merchant context. Name maximum 256 characters, region exactly 2 and mcc exactly 4; parent omitted when no child is available.
	MerchantInfo *model.EventMerchantInfo `json:"merchantInfo,omitempty"`
	// Optional approved context: at most 30 entries, keys at most 32 characters, values at most 128 and serialized JSON at most 3096. No key is approved in notification profile v2.1, so it is omitted.
	Metadata map[string]string `json:"metadata,omitempty"`
}
