package model

import "encoding/json"

// NullableApplyCardPaymentPreference preserves an explicit null preference in
// generated card-creation requests. A nil wrapper pointer omits the preference.
type NullableApplyCardPaymentPreference struct {
	value *ApplyCardPaymentPreference
	isSet bool
}

func (v NullableApplyCardPaymentPreference) Get() *ApplyCardPaymentPreference { return v.value }
func (v *NullableApplyCardPaymentPreference) Set(value *ApplyCardPaymentPreference) {
	v.value = value
	v.isSet = true
}
func (v NullableApplyCardPaymentPreference) IsSet() bool { return v.isSet }
func (v *NullableApplyCardPaymentPreference) Unset()     { v.value = nil; v.isSet = false }
func NewNullableApplyCardPaymentPreference(value *ApplyCardPaymentPreference) *NullableApplyCardPaymentPreference {
	return &NullableApplyCardPaymentPreference{value: value, isSet: true}
}
func (v NullableApplyCardPaymentPreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}
func (v *NullableApplyCardPaymentPreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
