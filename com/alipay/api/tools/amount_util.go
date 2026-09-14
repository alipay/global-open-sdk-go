package tools

import (
	"fmt"
	"strings"
)

const maxAmountValueLength = 16

// ToAmount converts an exact major-unit decimal string to Antom Amount.value.
func ToAmount(amount string, currency string) (string, error) {
	minorUnit, err := amountMinorUnit(currency)
	if err != nil {
		return "", err
	}
	point := -1
	for index := 0; index < len(amount); index++ {
		character := amount[index]
		if character == '.' && point < 0 && index > 0 && index < len(amount)-1 {
			point = index
		} else if character < '0' || character > '9' {
			return "", amountError("INVALID_AMOUNT_FORMAT", "amount must be an unsigned ASCII decimal string")
		}
	}
	if len(amount) == 0 {
		return "", amountError("INVALID_AMOUNT_FORMAT", "amount must be an unsigned ASCII decimal string")
	}
	whole, fraction := amount, ""
	if point >= 0 {
		whole, fraction = amount[:point], amount[point+1:]
	}
	if len(fraction) > minorUnit {
		if !allZeros(fraction[minorUnit:]) {
			return "", amountError("EXCESS_PRECISION", "amount exceeds the currency minor unit")
		}
		fraction = fraction[:minorUnit]
	}
	fraction += strings.Repeat("0", minorUnit-len(fraction))
	value := canonicalDigits(whole + fraction)
	if err := validateCanonical(value, currency); err != nil {
		return "", err
	}
	return value, nil
}

// FromAmount converts Antom Amount.value to a fixed-precision major-unit string.
func FromAmount(value string, currency string) (string, error) {
	minorUnit, err := amountMinorUnit(currency)
	if err != nil {
		return "", err
	}
	if err := validateValueFormat(value); err != nil {
		return "", err
	}
	canonical := canonicalDigits(value)
	if minorUnit == 0 {
		return canonical, nil
	}
	padded := canonical
	if len(canonical) <= minorUnit {
		padded = strings.Repeat("0", minorUnit+1-len(canonical)) + canonical
	}
	point := len(padded) - minorUnit
	return padded[:point] + "." + padded[point:], nil
}

// Validate checks an outbound Antom Amount.value against global currency rules.
func Validate(value string, currency string) error {
	if _, err := amountMinorUnit(currency); err != nil {
		return err
	}
	if err := validateValueFormat(value); err != nil {
		return err
	}
	return validateCanonical(value, currency)
}

func amountMinorUnit(currency string) (int, error) {
	if len(currency) != 3 || !upperASCII(currency[0]) || !upperASCII(currency[1]) || !upperASCII(currency[2]) {
		return 0, amountError("INVALID_CURRENCY", "currency must be three uppercase ASCII letters")
	}
	rules, err := loadAmountRules()
	if err != nil {
		return 0, err
	}
	rule, ok := rules.Currencies[currency]
	if !ok {
		return 0, amountError("UNSUPPORTED_CURRENCY", "currency is not supported by AmountUtil")
	}
	if rule.MinorUnit == nil {
		return 0, fmt.Errorf("RULE_DATA_ERROR: supported currency has no numeric minor unit")
	}
	return *rule.MinorUnit, nil
}

func validateValueFormat(value string) error {
	if len(value) == 0 {
		return amountError("INVALID_VALUE_FORMAT", "value must contain ASCII digits only")
	}
	for index := 0; index < len(value); index++ {
		if value[index] < '0' || value[index] > '9' {
			return amountError("INVALID_VALUE_FORMAT", "value must contain ASCII digits only")
		}
	}
	if len(value) > maxAmountValueLength {
		return amountError("VALUE_TOO_LONG", "value must contain at most 16 digits")
	}
	return nil
}

func validateCanonical(value string, currency string) error {
	if allZeros(value) {
		return amountError("AMOUNT_NOT_POSITIVE", "value must be greater than zero")
	}
	if len(value) > maxAmountValueLength {
		return amountError("VALUE_TOO_LONG", "value must contain at most 16 digits")
	}
	rules, err := loadAmountRules()
	if err != nil {
		return err
	}
	if constraint, ok := rules.AntomConstraints[currency]; ok {
		if !strings.HasSuffix(value, constraint.MinorValueMultiple[1:]) {
			return amountError("RULE_VIOLATION", "value does not satisfy the Antom currency constraint")
		}
	}
	return nil
}

func canonicalDigits(value string) string {
	canonical := strings.TrimLeft(value, "0")
	if canonical == "" {
		return "0"
	}
	return canonical
}

func allZeros(value string) bool {
	for index := 0; index < len(value); index++ {
		if value[index] != '0' {
			return false
		}
	}
	return true
}

func upperASCII(value byte) bool {
	return value >= 'A' && value <= 'Z'
}

func amountError(category string, detail string) error {
	return fmt.Errorf("%s: %s", category, detail)
}
