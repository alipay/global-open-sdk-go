package model

const (
	IOS     OsType = "IOS"
	ANDROID OsType = "ANDROID"
)

const (
	//BASE_USER_INFO, AGREEMENT_PAY, USER_INFO, USER_LOGIN_ID, HASH_LOGIN_ID, SEND_OTP;
	ScopeTypeBaseUserInfo ScopeType = "BASE_USER_INFO"
	ScopeTypeAgreementPay ScopeType = "AGREEMENT_PAY"
	ScopeTypeUserInfo     ScopeType = "USER_INFO"
	ScopeTypeUserLoginId  ScopeType = "USER_LOGIN_ID"
	ScopeTypeHashLoginId  ScopeType = "HASH_LOGIN_ID"
	ScopeTypeSendOtp      ScopeType = "SEND_OTP"
	ScopeTypeTAOBAOREBIND ScopeType = "TAOBAO_REBIND"
)

const (
	WEB      TerminalType = "WEB"
	WAP      TerminalType = "WAP"
	APP      TerminalType = "APP"
	MINI_APP TerminalType = "MINI_APP"
)

const (
	GrantTypeAUTHORIZATION_CODE GrantType = "AUTHORIZATION_CODE"
	GrantTypeREFRESH_TOKEN      GrantType = "REFRESH_TOKEN"
)

const (
	AGREEMENT_PAYMENT ProductCodeType = "AGREEMENT_PAYMENT"
	IN_STORE_PAYMENT  ProductCodeType = "IN_STORE_PAYMENT"
	CASHIER_PAYMENT   ProductCodeType = "CASHIER_PAYMENT"
)

const (
	BUNDLE  PresentmentMode = "BUNDLE"
	TILE    PresentmentMode = "TILE"
	UNIFIED PresentmentMode = "UNIFIED"
)

const (
	LegalEntityType_Company LegalEntityType = "COMPANY"
)

const (
	DisputeEvidenceType_DISPUTE_EVIDENCE_TEMPLATE DisputeEvidenceType = "DISPUTE_EVIDENCE_TEMPLATE"
	DisputeEvidenceType_DISPUTE_EVIDENCE_FILE     DisputeEvidenceType = "DISPUTE_EVIDENCE_FILE"
)
