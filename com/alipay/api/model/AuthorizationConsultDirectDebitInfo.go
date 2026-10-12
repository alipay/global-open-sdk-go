package model

type AuthorizationConsultDirectDebitInfo struct {
	ChannelProductCode  string                                  `json:"channelProductCode,omitempty"`
	PersonalProductCode string                                  `json:"personalProductCode,omitempty"`
	SignScene           string                                  `json:"signScene,omitempty"`
	AccessParams        *AuthorizationAccessParams              `json:"accessParams,omitempty"`
	PeriodRuleParams    *PeriodRuleParams                       `json:"periodRuleParams,omitempty"`
	PassBackParams      *AuthorizationConsultPassBackParams     `json:"passBackParams,omitempty"`
	SubscriptionParams  *AuthorizationConsultSubscriptionParams `json:"subscriptionParams,omitempty"`
	SubMerchant         *AuthorizationConsultSubMerchant        `json:"subMerchant,omitempty"`
	SubscribeType       string                                  `json:"subscribeType,omitempty"`
	OriAgreementNo      string                                  `json:"oriAgreementNo,omitempty"`
}
