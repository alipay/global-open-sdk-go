package model

type DisputeNotificationType string

const (
	DisputeNotificationType_DISPUTE_CREATED   DisputeNotificationType = "DISPUTE_CREATED"
	DisputeNotificationType_DISPUTE_JUDGED    DisputeNotificationType = "DISPUTE_JUDGED"
	DisputeNotificationType_DISPUTE_CANCELLED DisputeNotificationType = "DISPUTE_CANCELLED"
	DisputeNotificationType_DEFENSE_SUPPLIED  DisputeNotificationType = "DEFENSE_SUPPLIED"
	DisputeNotificationType_DEFENSE_DUE_ALERT DisputeNotificationType = "DEFENSE_DUE_ALERT"
	DisputeNotificationType_DISPUTE_ACCEPTED  DisputeNotificationType = "DISPUTE_ACCEPTED"
	RDR_RESOLVED                              DisputeNotificationType = "RDR_RESOLVED"
)

type DisputeJudgedResult string

const (
	DisputeJudgedResult_ACCEPT_BY_CUSTOMER DisputeJudgedResult = "ACCEPT_BY_CUSTOMER"
	DisputeJudgedResult_ACCEPT_BY_MERCHANT DisputeJudgedResult = "ACCEPT_BY_MERCHANT"
	DisputeJudgedResult_VALIDATE_SUCCESS   DisputeJudgedResult = "VALIDATE_SUCCESS"
	DisputeJudgedResult_VALIDATE_FAIL      DisputeJudgedResult = "VALIDATE_FAIL"
)

type DisputeAcceptReasonType string

const (
	MERCHANT_ACCEPTED          DisputeAcceptReasonType = "MERCHANT_ACCEPTED"
	TIMEOUT                    DisputeAcceptReasonType = "TIMEOUT"
	MANUAL_PROCESSING_ACCEPTED DisputeAcceptReasonType = "MANUAL_PROCESSING_ACCEPTED"
)

type DisputeType string

const (
	DisputeType_CHARGEBACK         DisputeType = "CHARGEBACK"
	DisputeType_RETRIEVAL_REQUEST  DisputeType = "RETRIEVAL_REQUEST"
	DisputeType_COMPLIANCE_REQUEST DisputeType = "COMPLIANCE_REQUEST"
)
