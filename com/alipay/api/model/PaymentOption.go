package model

type PaymentOption struct {
	PaymentMethodType     string                    `json:"paymentMethodType,omitempty"`
	PaymentMethodCategory PaymentMethodCategoryType `json:"paymentMethodCategory,omitempty"`
	PaymentMethodRegion   []string                  `json:"paymentMethodRegion,omitempty"`
	Enabled               bool                      `json:"enabled,omitempty"`
	Preferred             bool                      `json:"preferred,omitempty"`
	DisableReason         string                    `json:"disableReason,omitempty"`
	AmountLimitInfoMap    *map[string]AmountLimit   `json:"amountLimitInfoMap,omitempty"`
	SupportedCurrencies   []string                  `json:"supportedCurrencies,omitempty"`
	PaymentOptionDetail   *PaymentOptionDetail      `json:"paymentOptionDetail,omitempty"`
	ExtendInfo            string                    `json:"extendInfo,omitempty"`
	Logo                  *Logo                     `json:"logo,omitempty"`
	PromoNames            []string                  `json:"promoNames,omitempty"`
	Installment           *Installment              `json:"installment,omitempty"`
	PromotionInfos        []*PromotionInfo          `json:"promotionInfos,omitempty"`
}
