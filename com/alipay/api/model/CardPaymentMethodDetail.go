package model

type CardPaymentMethodDetail struct {
	SupportedBrands             string    `json:"supportedBrands,omitempty"`
	CardToken                   string    `json:"cardToken,omitempty"`
	CardNo                      string    `json:"cardNo,omitempty"`
	Brand                       CardBrand `json:"brand,omitempty"`
	SelectedCardBrand           CardBrand `json:"selectedCardBrand,omitempty"`
	CardIssuer                  string    `json:"cardIssuer,omitempty"`
	CountryIssue                string    `json:"countryIssue,omitempty"`
	InstUserName                *UserName `json:"instUserName,omitempty"`
	ExpiryYear                  string    `json:"expiryYear,omitempty"`
	ExpiryMonth                 string    `json:"expiryMonth,omitempty"`
	BillingAddress              *Address  `json:"billingAddress,omitempty"`
	Mask                        string    `json:"mask,omitempty"`
	Last4                       string    `json:"last4,omitempty"`
	PaymentMethodDetailMetadata string    `json:"paymentMethodDetailMetadata,omitempty"`
	MaskedCardNo                string    `json:"maskedCardNo,omitempty"`
	Fingerprint                 string    `json:"fingerprint,omitempty"`
	AuthenticationFlow          string    `json:"authenticationFlow,omitempty"`
	Funding                     string    `json:"funding,omitempty"`
	AvsResultRaw                string    `json:"avsResultRaw,omitempty"`
	CvvResultRaw                string    `json:"cvvResultRaw,omitempty"`
	Bin                         string    `json:"bin,omitempty"`
	IssuerName                  string    `json:"issuerName,omitempty"`
	IssuingCountry              string    `json:"issuingCountry,omitempty"`
	LastFour                    string    `json:"lastFour,omitempty"`
	CardholderName              *UserName `json:"cardholderName,omitempty"`
	Cvv                         string    `json:"cvv,omitempty"`
	DateOfBirth                 string    `json:"dateOfBirth,omitempty"`
	BusinessNo                  string    `json:"businessNo,omitempty"`
	CardPasswordDigest          string    `json:"cardPasswordDigest,omitempty"`
	Cpf                         string    `json:"cpf,omitempty"`
	PayerEmail                  string    `json:"payerEmail,omitempty"`
	NetworkTransactionId        string    `json:"networkTransactionId,omitempty"`
	Is3DSAuthentication         bool      `json:"is3DSAuthentication,omitempty"`
	Request3DS                  string    `json:"request3DS,omitempty"`
	ScaExemptionIndicator       string    `json:"scaExemptionIndicator,omitempty"`
	EnableAuthenticationUpgrade string    `json:"enableAuthenticationUpgrade,omitempty"`
	MpiData                     *MpiData  `json:"mpiData,omitempty"`
}
