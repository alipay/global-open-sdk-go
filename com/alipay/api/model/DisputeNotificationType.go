package model

type DisputeNotificationType string

const (
	DisputeNotificationType_DISPUTE_CREATED   DisputeNotificationType = "DISPUTE_CREATED"
	DisputeNotificationType_DISPUTE_JUDGED    DisputeNotificationType = "DISPUTE_JUDGED"
	DisputeNotificationType_DISPUTE_CANCELLED DisputeEvidenceType     = "DISPUTE_CANCELLED"
	DisputeNotificationType_DEFENSE_SUPPLIED  DisputeEvidenceType     = "DEFENSE_SUPPLIED"
	DisputeNotificationType_DEFENSE_DUE_ALERT DisputeEvidenceType     = "DEFENSE_DUE_ALERT"
	DisputeNotificationType_DISPUTE_ACCEPTED  DisputeEvidenceType     = "DISPUTE_ACCEPTED"
)

type DisputeJudgedResult string

const (
	DisputeJudgedResult_ACCEPT_BY_CUSTOMER DisputeJudgedResult = "ACCEPT_BY_CUSTOMER"
	DisputeJudgedResult_ACCEPT_BY_MERCHANT DisputeJudgedResult = "ACCEPT_BY_MERCHANT"
)
