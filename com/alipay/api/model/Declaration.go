package model

type Declaration struct {
	DeclarationBizScene      DeclarationBizSceneType `json:"declarationBizScene,omitempty"`
	DeclarationBeneficiaryId string                  `json:"declarationBeneficiaryId,omitempty"`
}
