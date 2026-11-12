package model

type TradeBillQueryInfo struct {
	MerchantId            string                  `json:"merchantId,omitempty"`
	ReferenceStoreId      string                  `json:"referenceStoreId,omitempty"`
	Tid                   string                  `json:"tid,omitempty"`
	MerchantTransactionId string                  `json:"merchantTransactionId,omitempty"`
	PaymentAmount         *TradeBillPaymentAmount `json:"paymentAmount,omitempty"`
	BizOrderType          string                  `json:"bizOrderType,omitempty"`
	BizOrderStatus        string                  `json:"bizOrderStatus,omitempty"`
	StartTime             string                  `json:"startTime,omitempty"`
	CompletedTime         string                  `json:"completedTime,omitempty"`
}
