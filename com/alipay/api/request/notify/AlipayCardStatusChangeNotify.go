// notify/card_status_change_notify.go
package notify

// AlipayCardStatusChangeNotify 卡状态变更通知
type AlipayCardStatusChangeNotify struct {
	AlipayNotify

	RequestID   string `json:"requestId"`
	AssetId     string `json:"assetId"`
	OperateType string `json:"operateType"`         // ADD / FREEZE / UNFREEZE / CANCEL
	Status      string `json:"status"`              // SUCCESS / FAIL
	CardStatus  string `json:"cardStatus"`          // ACTIVE / FROZEN / CANCEL
	CreatedTime string `json:"createdTime"`         // ISO 8601
	UpdatedTime string `json:"updatedTime"`         // ISO 8601
	CardBrand   string `json:"cardBrand,omitempty"` // e.g., "MASTERCARD"
	CardType    string `json:"cardType,omitempty"`  // VIRTUAL / PHYSICAL
}
