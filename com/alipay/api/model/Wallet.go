package model

type Wallet struct {
	AccountNo         string    `json:"accountNo,omitempty"`
	AccountHolderName *UserName `json:"accountHolderName,omitempty"`
	PhoneNo           string    `json:"phoneNo,omitempty"`
	Email             string    `json:"email,omitempty"`
	BillingAddress    *Address  `json:"billingAddress,omitempty"`
}
