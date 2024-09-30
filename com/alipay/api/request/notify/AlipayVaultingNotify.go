package notify

type AlipayVaultingNotify struct {
	AlipayNotify
	VaultingRequestId   string `json:"vaultingRequestId,omitempty"`
	PaymentMethodDetail string `json:"paymentMethodDetail,omitempty"`
	VaultingCreateTime  string `json:"vaultingCreateTime,omitempty"`
	AcquirerInfo        string `json:"acquirerInfo,omitempty"`
}
