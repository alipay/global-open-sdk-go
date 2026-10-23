package notify

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDefenseNotifications(t *testing.T) {
	strategy := `{"disputeId":"d1","notifyTime":"2026-10-22T01:02:03.123+08:00","winRateLevel":"FUTURE","evidenceList":[{"evidenceTag":"FUTURE_TAG","obtained":false,"winRateContribution":"FUTURE"},{"evidenceTag":"ORDER_INFORMATION","obtained":true}],"cardholderDisputeReason":"Reason","defenseStrategyDescription":"Supply evidence","futureField":{"ignored":true}}`
	document := `{"disputeId":"d1","notifyTime":"2026-10-22T01:02:03.123+08:00","defenseFile":{"fileName":"defense.pdf","fileFormat":"FUTURE","fileMd5":"5d41402abc4b2a76b9719d911017c592"},"winRateLevel":"FUTURE","recommendedEvidence":{"FUTURE_TAG":"提供材料","ORDER_INFORMATION":"Order proof"},"suggestionsForImprovement":"Add proof"}`
	var s AlipayDefenseStrategySuggestionNotify
	if err := json.Unmarshal([]byte(strategy), &s); err != nil {
		t.Fatal(err)
	}
	if s.EvidenceList[0].Obtained == nil || *s.EvidenceList[0].Obtained || !*s.EvidenceList[1].Obtained {
		t.Fatal("boolean values lost")
	}
	var d AlipayDefenseDocumentNotify
	if err := json.Unmarshal([]byte(document), &d); err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		input string
		value interface{}
	}{{`{"disputeId":"d1","notifyTime":"2026-10-22T01:02:03.123+08:00","winRateLevel":"FUTURE","evidenceList":[{"evidenceTag":"FUTURE_TAG","obtained":false,"winRateContribution":"FUTURE"},{"evidenceTag":"ORDER_INFORMATION","obtained":true}],"cardholderDisputeReason":"Reason","defenseStrategyDescription":"Supply evidence"}`, s}, {document, d}} {
		out, err := json.Marshal(c.value)
		if err != nil {
			t.Fatal(err)
		}
		var want, got interface{}
		if err := json.Unmarshal([]byte(c.input), &want); err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(out, &got); err != nil {
			t.Fatal(err)
		}
		// AlipayNotify embeds a non-pointer Result; its zero value is serialized.
		// Preserve that existing base behavior while checking all new wire fields.
		want.(map[string]interface{})["result"] = map[string]interface{}{
			"resultCode": "", "resultStatus": "", "resultMessage": "",
		}
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wire fields changed: %s", out)
		}
	}
}

func TestDefenseNotificationMissingOptionals(t *testing.T) {
	for _, body := range []string{`{}`, `{"winRateLevel":null,"recommendedEvidence":null,"suggestionsForImprovement":null}`} {
		var d AlipayDefenseDocumentNotify
		if err := json.Unmarshal([]byte(body), &d); err != nil {
			t.Fatal(err)
		}
		if d.WinRateLevel != "" || d.RecommendedEvidence != nil || d.SuggestionsForImprovement != "" {
			t.Fatal("optional fields changed")
		}
	}
	var s AlipayDefenseStrategySuggestionNotify
	if err := json.Unmarshal([]byte(`{"evidenceList":[{"evidenceTag":"X"}]}`), &s); err != nil {
		t.Fatal(err)
	}
	if s.EvidenceList[0].Obtained != nil {
		t.Fatal("missing boolean must be nil")
	}
}
