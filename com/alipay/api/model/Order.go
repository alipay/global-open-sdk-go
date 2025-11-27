package model

type Order struct {
	ReferenceOrderId    string       `json:"referenceOrderId,omitempty"`
	OrderDescription    string       `json:"orderDescription,omitempty"`
	OrderAmount         *Amount      `json:"orderAmount,omitempty"`
	OrderDiscountAmount *Amount      `json:"orderDiscountAmount,omitempty"`
	SubTotalOrderAmount *Amount      `json:"subTotalOrderAmount,omitempty"`
	Merchant            *Merchant    `json:"merchant,omitempty"`
	Goods               []*Goods     `json:"goods,omitempty"`
	Shipping            *Shipping    `json:"shipping,omitempty"`
	Buyer               *Buyer       `json:"buyer,omitempty"`
	Env                 *Env         `json:"env,omitempty"`
	ExtendInfo          string       `json:"extendInfo,omitempty"`
	Transit             *Transit     `json:"transit,omitempty"`
	Lodging             *Lodging     `json:"lodging,omitempty"`
	Gaming              *Gaming      `json:"gaming,omitempty"`
	NeedDeclaration     bool         `json:"needDeclaration,omitempty"`
	Declaration         *Declaration `json:"declaration,omitempty"`
	OrderType           string       `json:"orderType,omitempty"`
}
