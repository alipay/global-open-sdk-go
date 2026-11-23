package model

import (
	"encoding/json"
	"testing"
)

func TestPaymentPreferenceJSON(t *testing.T) {
	cases := []struct {
		name  string
		order []string
		want  string
	}{
		{"inherit", nil, `{}`},
		{"explicit empty", []string{}, `{"currencyOrder":[]}`},
		{"ordered custom", []string{"HKD", "USD"}, `{"currencyOrder":["HKD","USD"]}`},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			preference := PaymentPreference{CurrencyOrder: test.order}
			got, err := json.Marshal(preference)
			if err != nil || string(got) != test.want {
				t.Fatalf("MarshalJSON = %s, %v; want %s", got, err, test.want)
			}
			var parsed PaymentPreference
			if err := json.Unmarshal(got, &parsed); err != nil {
				t.Fatal(err)
			}
			if (parsed.CurrencyOrder == nil) != (test.order == nil) {
				t.Fatal("round trip changed omitted versus explicit empty order")
			}
			if len(parsed.CurrencyOrder) != len(test.order) {
				t.Fatal("round trip changed order length")
			}
			for i := range test.order {
				if parsed.CurrencyOrder[i] != test.order[i] {
					t.Fatal("round trip changed currency priority")
				}
			}
		})
	}
}
