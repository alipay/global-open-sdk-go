package model

type Certificate struct {
	CertificateType      CertificateType `json:"certificateType,omitempty"`
	CertificateNo        string          `json:"certificateNo,omitempty"`
	HolderName           *UserName       `json:"holderName,omitempty"`
	FileKeys             []string        `json:"fileKeys,omitempty"`
	CertificateAuthority string          `json:"certificateAuthority,omitempty"`
}
