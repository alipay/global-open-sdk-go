package model

type CompanyType string

const (
	CompanyType_ENTERPRISE               CompanyType = "ENTERPRISE"
	CompanyType_PARTNERSHIP              CompanyType = "PARTNERSHIP"
	CompanyType_SOLE_PROPRIETORSHIP      CompanyType = "SOLE_PROPRIETORSHIP"
	CompanyType_STATE_OWNED_BUSINESS     CompanyType = "STATE_OWNED_BUSINESS"
	CompanyType_PRIVATELY_OWNED_BUSINESS CompanyType = "PRIVATELY_OWNED_BUSINESS"
	CompanyType_PUBLICLY_LISTED_BUSINESS CompanyType = "PUBLICLY_LISTED_BUSINESS"
	CompanyType_LTDA                     CompanyType = "LTDA"
	CompanyType_SA                       CompanyType = "SA"
	CompanyType_EIRELI                   CompanyType = "EIRELI"
	CompanyType_BOFC                     CompanyType = "BOFC"
	CompanyType_MEI                      CompanyType = "MEI"
	CompanyType_EI                       CompanyType = "EI"
)
