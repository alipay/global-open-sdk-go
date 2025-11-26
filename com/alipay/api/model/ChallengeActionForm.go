package model

type ChallengeActionForm struct {
	ChallengeType        ChallengeType              `json:"challengeType,omitempty"`
	ChallengeRenderValue string                     `json:"challengeRenderValue,omitempty"`
	TriggerSource        ChallengeTriggerSourceType `json:"triggerSource,omitempty"`
	ExtendInfo           string                     `json:"extendInfo,omitempty"`
}
