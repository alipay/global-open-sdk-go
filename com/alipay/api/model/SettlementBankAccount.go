package model

type SettlementBankAccount struct {
	BankAccountNo        string            `json:"bankAccountNo,omitempty"`
	AccountHolderName    string            `json:"accountHolderName,omitempty"`
	SwiftCode            string            `json:"swiftCode,omitempty"`
	BankRegion           string            `json:"bankRegion,omitempty"`
	AccountHolderType    AccountHolderType `json:"accountHolderType,omitempty"`
	RoutingNumber        string            `json:"routingNumber,omitempty"`
	BranchCode           string            `json:"branchCode,omitempty"`
	AccountHolderTIN     string            `json:"accountHolderTIN,omitempty"`
	AccountType          AccountType       `json:"accountType,omitempty"`
	BankName             string            `json:"bankName,omitempty"`
	AccountHolderAddress *Address          `json:"accountHolderAddress,omitempty"`
	Iban                 string            `json:"iban,omitempty"`
}
