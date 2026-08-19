package tools

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"
)

//go:embed amount-currency-rules.json
var amountRulesJSON []byte

type amountCurrencyRule struct {
	MinorUnit *int `json:"minorUnit"`
}

type amountConstraint struct {
	MinorValueMultiple string `json:"minorValueMultiple"`
}

type amountRuleSet struct {
	SchemaVersion    int                           `json:"schemaVersion"`
	Currencies       map[string]amountCurrencyRule `json:"currencies"`
	AntomConstraints map[string]amountConstraint   `json:"antomConstraints"`
}

var (
	loadedAmountRules amountRuleSet
	amountRulesError  error
	amountRulesOnce   sync.Once
)

func loadAmountRules() (*amountRuleSet, error) {
	amountRulesOnce.Do(func() {
		if err := json.Unmarshal(amountRulesJSON, &loadedAmountRules); err != nil {
			amountRulesError = fmt.Errorf("RULE_DATA_ERROR: unable to parse amount currency rules: %w", err)
			return
		}
		if loadedAmountRules.SchemaVersion != 1 || loadedAmountRules.Currencies == nil || loadedAmountRules.AntomConstraints == nil {
			amountRulesError = fmt.Errorf("RULE_DATA_ERROR: invalid amount currency rules structure")
			return
		}
		for _, rule := range loadedAmountRules.Currencies {
			if rule.MinorUnit != nil && (*rule.MinorUnit < 0 || *rule.MinorUnit > 4) {
				amountRulesError = fmt.Errorf("RULE_DATA_ERROR: invalid minor unit")
				return
			}
		}
		for _, constraint := range loadedAmountRules.AntomConstraints {
			if !isPowerOfTen(constraint.MinorValueMultiple) {
				amountRulesError = fmt.Errorf("RULE_DATA_ERROR: invalid amount constraint")
				return
			}
		}
	})
	if amountRulesError != nil {
		return nil, amountRulesError
	}
	return &loadedAmountRules, nil
}

func isPowerOfTen(value string) bool {
	if len(value) == 0 || value[0] != '1' {
		return false
	}
	return allZeros(value[1:])
}
