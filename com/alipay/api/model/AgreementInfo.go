package model

type AgreementInfo struct {
	AuthState          string `json:"authState,omitempty"`
	UserLoginId        string `json:"userLoginId,omitempty"`
	UserLoginType      string `json:"userLoginType,omitempty"`
	DisplayUserLoginId string `json:"displayUserLoginId,omitempty"`
}
