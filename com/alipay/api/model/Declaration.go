package model

type Declaration struct {
	DeclarationBizScene      DeclarationBizSceneType `json:"declarationBizScene,omitempty"`
	DeclarationBeneficiaryId string                  `json:"declarationBeneficiaryId,omitempty"`
}

type DeclarationBizSceneType string

const (
	DeclarationBizSceneType_AIRLINE DeclarationBizSceneType = "AIRLINE"
	DeclarationBizSceneType_HOTEL   DeclarationBizSceneType = "HOTEL"
)
