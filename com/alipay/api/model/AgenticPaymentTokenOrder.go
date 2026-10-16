package model

type AgenticPaymentTokenOrder struct {
	PaymentAmount    *AgenticAmount                    `json:"paymentAmount,omitempty"`
	Merchant         *AgenticPaymentTokenOrderMerchant `json:"merchant,omitempty"`
	Goods            []*AgenticPaymentTokenOrderGoods  `json:"goods,omitempty"`
	ReferenceOrderId string                            `json:"referenceOrderId,omitempty"`
}
