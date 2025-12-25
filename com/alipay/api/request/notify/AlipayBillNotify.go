// notify/bill_notify.go
package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

// AlipayBillNotify 交易账单状态通知
type AlipayBillNotify struct {
	AlipayNotify

	AssetId         string `json:"assetId"`
	MaskedCardNo    string `json:"maskedCardNo"`
	OrderNo         string `json:"orderNo"`
	CardNickName    string `json:"cardNickName,omitempty"`
	TransactionTime string `json:"transactionTime"`
	MerchantName    string `json:"merchantName,omitempty"`

	TradeAmount *model.Amount `json:"tradeAmount"`
	InAmount    *model.Amount `json:"inAmount,omitempty"`
	OutAmount   *model.Amount `json:"outAmount,omitempty"`

	ExchangeRate string `json:"exchangeRate,omitempty"`
	BillType     string `json:"billType"`               // e.g., "CARD_PAYMENT"
	TradeCountry string `json:"tradeCountry,omitempty"` // ISO-3166 2-letter
	BillStatus   string `json:"billStatus"`             // e.g., "SUCCESS", "FAILED"

	BillFailReason *FailReason `json:"billFailReason,omitempty"`

	LastUpdate string            `json:"lastUpdate"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// FailReason 交易失败原因
type FailReason struct {
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorDesc string `json:"errorDesc,omitempty"`
}
