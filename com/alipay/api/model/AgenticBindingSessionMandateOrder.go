package model

type AgenticBindingSessionMandateOrder struct {
	PaymentAmount    *AgenticAmount                             `json:"paymentAmount,omitempty"`
	Merchant         *AgenticBindingSessionMandateOrderMerchant `json:"merchant,omitempty"`
	Goods            []*AgenticSessionGoods                     `json:"goods,omitempty"`
	CheckoutJwt      string                                     `json:"checkoutJwt,omitempty"`
	ReferenceOrderId string                                     `json:"referenceOrderId,omitempty"`
}
