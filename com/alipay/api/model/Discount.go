package model

type Discount struct {
	DiscountName  string `json:"discountName,omitempty"`
	SavingsAmount Amount `json:"savingsAmount,omitempty"`
}
