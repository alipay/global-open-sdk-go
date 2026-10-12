package model

type AuthorizationConsultSubscriptionParams struct {
	SubscribeTplType        string `json:"subscribeTplType,omitempty"`
	FirstSubscriptionInfo   string `json:"firstSubscriptionInfo,omitempty"`
	FollowSubscriptionInfos string `json:"followSubscriptionInfos,omitempty"`
	OriginSubscriptionInfo  string `json:"originSubscriptionInfo,omitempty"`
	TargetSubscriptionInfo  string `json:"targetSubscriptionInfo,omitempty"`
	NextPaymentDate         string `json:"nextPaymentDate,omitempty"`
	ExpireAt                string `json:"expireAt,omitempty"`
	BaseAmount              string `json:"baseAmount,omitempty"`
	OffsetAmount            string `json:"offsetAmount,omitempty"`
	DeductName              string `json:"deductName,omitempty"`
	DeductDesc              string `json:"deductDesc,omitempty"`
}
