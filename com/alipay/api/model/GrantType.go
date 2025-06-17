package model

type GrantType string

const (
	GrantType_AUTHORIZATION_CODE GrantType = "AUTHORIZATION_CODE"
	GrantType_REFRESH_TOKEN      GrantType = "REFRESH_TOKEN"
)
