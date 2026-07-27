package model

type Meter struct {
	MeterId           string `json:"meterId,omitempty"`
	MeterName         string `json:"meterName,omitempty"`
	EventName         string `json:"eventName,omitempty"`
	Status            string `json:"status,omitempty"`
	AggregationMethod string `json:"aggregationMethod,omitempty"`
	EventTimeWindow   string `json:"eventTimeWindow,omitempty"`
	ValueKeyOverride  string `json:"valueKeyOverride,omitempty"`
	CreatedDateTime   string `json:"createdDateTime,omitempty"`
	UpdatedDateTime   string `json:"updatedDateTime,omitempty"`
}
