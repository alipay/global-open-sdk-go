package model

import (

)




type TrialPlan struct {
        TrialEndTime string `json:"trialEndTime,omitempty"`
        FreeTrialDays int32 `json:"freeTrialDays,omitempty"`
        Phases [] *TrialPhase `json:"phases,omitempty"` 
}








