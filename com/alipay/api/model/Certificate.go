package model

type Certificate struct {
	CertificateType      CertificateType `json:"certificate_type,omitempty"`
	CertificateNo        string          `json:"certificate_no,omitempty"`
	HolderName           *UserName       `json:"holder_name,omitempty"`
	FileKeys             []string        `json:"file_keys,omitempty"`
	CertificateAuthority string          `json:"certificate_authority,omitempty"`
	GrantType            string          `json:"grant_type,omitempty"`
}

type CertificateType string

const (
	//ENTERPRISE_REGISTRATION,
	//LICENSE_INFO,
	//ID_CARD,
	//PASSPORT,
	//DRIVING_LICENSE,
	//CPF,
	//CNPJ,

	CertificateType_ENTERPRISE_REGISTRATION CertificateType = "ENTERPRISE_REGISTRATION"
	CertificateType_LICENSE_INFO            CertificateType = "LICENSE_INFO"
	CertificateType_ID_CARD                 CertificateType = "ID_CARD"
	CertificateType_PASSPORT                CertificateType = "PASSPORT"
	CertificateType_DRIVING_LICENSE         CertificateType = "DRIVING_LICENSE"
	CertificateType_CPF                     CertificateType = "CPF"
	CertificateType_CNPJ                    CertificateType = "CNPJ"
)
