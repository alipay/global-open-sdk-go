package model

type ChallengeType string

const (
	ChallengeType_SMS_OTP           ChallengeType = "SMS_OTP"
	ChallengeType_PLAINTEXT_CARD_NO ChallengeType = "PLAINTEXT_CARD_NO"
	ChallengeType_CARD_EXPIRE_DATE  ChallengeType = "CARD_EXPIRE_DATE"
)
