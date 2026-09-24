package model

type PaymentFactor struct {
	StorePaymentMethodForBuyer bool                   `json:"storePaymentMethodForBuyer,omitempty"`
	IsPaymentEvaluation        bool                   `json:"isPaymentEvaluation,omitempty"`
	InStorePaymentScenario     InStorePaymentScenario `json:"inStorePaymentScenario,omitempty"`
	PresentmentMode            PresentmentMode        `json:"presentmentMode,omitempty"`
	CaptureMode                string                 `json:"captureMode,omitempty"`
	IsAuthorization            bool                   `json:"isAuthorization,omitempty"`
}
