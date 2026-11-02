package notify

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestCardTransactionEventNotifyFixtures(t *testing.T) {
	raw, err := os.ReadFile("testdata/card-transaction-event-notify.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixtures []struct {
		Name string
		Body json.RawMessage
	}
	if err := json.Unmarshal(raw, &fixtures); err != nil {
		t.Fatal(err)
	}
	if reflect.TypeOf(AlipayCardTransactionEventNotify{}).NumField() != 22 {
		t.Fatal("expected exactly 22 flat event fields")
	}
	for _, fixture := range fixtures {
		t.Run(fixture.Name, func(t *testing.T) {
			var expected map[string]any
			if err := json.Unmarshal(fixture.Body, &expected); err != nil {
				t.Fatal(err)
			}
			withUnknown := make(map[string]any)
			for key, value := range expected {
				withUnknown[key] = value
			}
			withUnknown["futureField"] = map[string]any{"ignored": true}
			body, err := json.Marshal(withUnknown)
			if err != nil {
				t.Fatal(err)
			}
			var notify AlipayCardTransactionEventNotify
			if err := json.Unmarshal(body, &notify); err != nil {
				t.Fatal(err)
			}
			encoded, err := json.Marshal(notify)
			if err != nil {
				t.Fatal(err)
			}
			var actual map[string]any
			if err := json.Unmarshal(encoded, &actual); err != nil {
				t.Fatal(err)
			}
			for key, value := range expected {
				if value == nil {
					delete(expected, key)
				}
			}
			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("wire fields changed\nexpected %v\nactual %v", expected, actual)
			}
			if fixture.Name == "capture" {
				if notify.OutAmount == nil || notify.OutAmount.Value != "0" ||
					notify.OutAmount.Currency != "USD" || notify.InAmount != nil ||
					notify.TransactionAmount.Value != "30000" ||
					notify.AuthTime != "2026-07-09T10:35:10+00:00" ||
					notify.BalanceType != "ANTOM_BIZ_ACCOUNT" {
					t.Fatal("CAPTURE mapping changed")
				}
			}
			if fixture.Name == "auth" && notify.AuthTime != "" {
				t.Fatal("AUTH authTime was inferred")
			}
		})
	}
}
