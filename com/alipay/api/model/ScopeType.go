package model

type ScopeType string

const (
	ScopeType_BASE_USER_INFO ScopeType = "BASE_USER_INFO"
	ScopeType_AGREEMENT_PAY  ScopeType = "AGREEMENT_PAY"
	ScopeType_USER_INFO      ScopeType = "USER_INFO"
	ScopeType_USER_LOGIN_ID  ScopeType = "USER_LOGIN_ID"
	ScopeType_HASH_LOGIN_ID  ScopeType = "HASH_LOGIN_ID"
	ScopeType_SEND_OTP       ScopeType = "SEND_OTP"
	ScopeType_TAOBAO_REBIND  ScopeType = "TAOBAO_REBIND"
)
