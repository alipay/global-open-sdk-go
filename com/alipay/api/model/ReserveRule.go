package model

type ReserveRule struct {
	RuleId             string             `json:"ruleId,omitempty"`
	MerchantId         string             `json:"merchantId,omitempty"`
	FundsType          FundsType          `json:"fundsType,omitempty"`
	TakeType           TakeType           `json:"takeType,omitempty"`
	PaymentMethodScope PaymentMethodScope `json:"paymentMethodScope,omitempty"`
	Ratio              int32              `json:"ratio,omitempty"`
	ReleaseType        ReleaseType        `json:"releaseType,omitempty"`
	ReleaseTime        string             `json:"releaseTime,omitempty"`
	RetentionTime      int32              `json:"retentionTime,omitempty"`
	RuleStatus         RuleStatus         `json:"ruleStatus,omitempty"`
	CreateTime         string             `json:"createTime,omitempty"`
	UpdateTime         string             `json:"updateTime,omitempty"`
}
