package model

type CertificateType string

const (
	CertificateType_ENTERPRISE_REGISTRATION CertificateType = "ENTERPRISE_REGISTRATION"
	CertificateType_LICENSE_INFO            CertificateType = "LICENSE_INFO"
	CertificateType_ID_CARD                 CertificateType = "ID_CARD"
	CertificateType_PASSPORT                CertificateType = "PASSPORT"
	CertificateType_DRIVING_LICENSE         CertificateType = "DRIVING_LICENSE"
	CertificateType_CPF                     CertificateType = "CPF"
	CertificateType_CNPJ                    CertificateType = "CNPJ"
)
