package model

type FundingType string

const (
	FundingType_CREDIT         FundingType = "CREDIT"
	FundingType_DEBIT          FundingType = "DEBIT"
	FundingType_PREPAID        FundingType = "PREPAID"
	FundingType_CHARGE         FundingType = "CHARGE"
	FundingType_DEFERRED_DEBIT FundingType = "DEFERRED_DEBIT"
)
