package tools

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

type amountVector struct {
	Name     string `json:"name"`
	Amount   string `json:"amount"`
	Value    string `json:"value"`
	Currency string `json:"currency"`
	Result   string `json:"result"`
	Error    string `json:"error"`
}

type amountVectors struct {
	ToAmount   []amountVector `json:"toAmount"`
	FromAmount []amountVector `json:"fromAmount"`
	Validate   []amountVector `json:"validate"`
}

func TestAmountUtilSharedVectors(t *testing.T) {
	data, err := os.ReadFile("testdata/amount-conversion-test-cases.json")
	if err != nil {
		t.Fatal(err)
	}
	var vectors amountVectors
	if err := json.Unmarshal(data, &vectors); err != nil {
		t.Fatal(err)
	}
	for _, vector := range vectors.ToAmount {
		vector := vector
		t.Run("to/"+vector.Name, func(t *testing.T) {
			result, err := ToAmount(vector.Amount, vector.Currency)
			checkAmountVector(t, vector, result, err)
		})
		if vector.Result != "" {
			major, err := FromAmount(vector.Result, vector.Currency)
			if err != nil {
				t.Fatal(err)
			}
			value, err := ToAmount(major, vector.Currency)
			if err != nil || value != vector.Result {
				t.Fatalf("%s round trip: value=%s error=%v", vector.Name, value, err)
			}
		}
	}
	for _, vector := range vectors.FromAmount {
		vector := vector
		t.Run("from/"+vector.Name, func(t *testing.T) {
			result, err := FromAmount(vector.Value, vector.Currency)
			checkAmountVector(t, vector, result, err)
		})
	}
	for _, vector := range vectors.Validate {
		vector := vector
		t.Run("validate/"+vector.Name, func(t *testing.T) {
			checkAmountVector(t, vector, "", Validate(vector.Value, vector.Currency))
		})
	}
}

func checkAmountVector(t *testing.T, vector amountVector, result string, err error) {
	t.Helper()
	if vector.Error != "" {
		if err == nil || !strings.HasPrefix(err.Error(), vector.Error+":") {
			t.Fatalf("expected %s, got %v", vector.Error, err)
		}
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if vector.Result != "" && result != vector.Result {
		t.Fatalf("expected %s, got %s", vector.Result, result)
	}
}
