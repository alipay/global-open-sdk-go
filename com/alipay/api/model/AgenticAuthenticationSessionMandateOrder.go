package model

type AgenticAuthenticationSessionMandateOrder struct {
	PaymentAmount    *AgenticAmount                                    `json:"paymentAmount,omitempty"`
	Merchant         *AgenticAuthenticationSessionMandateOrderMerchant `json:"merchant,omitempty"`
	Goods            []*AgenticSessionGoods                            `json:"goods,omitempty"`
	CheckoutJwt      string                                            `json:"checkoutJwt,omitempty"`
	ReferenceOrderId string                                            `json:"referenceOrderId,omitempty"`
}
