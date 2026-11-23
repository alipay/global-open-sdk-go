package model

import "encoding/json"

// MarshalJSON preserves an explicit empty order for server-side validation.
// A nil order remains omitted so the empty preference object can inherit policy.
func (p PaymentPreference) MarshalJSON() ([]byte, error) {
	type plain PaymentPreference
	var order *[]string
	if p.CurrencyOrder != nil {
		order = &p.CurrencyOrder
	}
	return json.Marshal(struct {
		*plain
		CurrencyOrder *[]string `json:"currencyOrder,omitempty"`
	}{plain: (*plain)(&p), CurrencyOrder: order})
}
