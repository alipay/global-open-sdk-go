package notify

import (
	"encoding/json"
	"testing"
)

func TestDisputeEventTimes(t *testing.T) {
	cases := []struct{ event, key, value string }{
		{"DISPUTE_CANCELLED", "disputeCancelledTime", "2026-08-03T03:56:00+00:00"},
		{"DEFENSE_SUPPLIED", "defenseSuppliedTime", "2022-09-22T01:10:05-07:00"},
	}
	for _, tc := range cases {
		t.Run(tc.event, func(t *testing.T) {
			body, _ := json.Marshal(map[string]string{"disputeId": "d1", "disputeNotificationType": tc.event, tc.key: tc.value})
			var n AlipayDisputeNotify
			if err := json.Unmarshal(body, &n); err != nil {
				t.Fatal(err)
			}
			got, other := n.DisputeCancelledTime, n.DefenseSuppliedTime
			if tc.key == "defenseSuppliedTime" {
				got, other = other, got
			}
			if got != tc.value || other != "" || n.DisputeId != "d1" {
				t.Fatalf("unexpected notification: %+v", n)
			}
			encoded, err := json.Marshal(n)
			if err != nil {
				t.Fatal(err)
			}
			var out map[string]interface{}
			if err := json.Unmarshal(encoded, &out); err != nil {
				t.Fatal(err)
			}
			if out[tc.key] != tc.value {
				t.Fatalf("wire key/time changed: %s", encoded)
			}
		})
	}
}

func TestDisputeLegacyTimesOmitted(t *testing.T) {
	for _, body := range []string{`{"disputeId":"legacy","disputeTime":"2022-09-20T23:41:32-07:00"}`, `{"disputeCancelledTime":null,"defenseSuppliedTime":null}`} {
		var n AlipayDisputeNotify
		if err := json.Unmarshal([]byte(body), &n); err != nil {
			t.Fatal(err)
		}
		if n.DisputeCancelledTime != "" || n.DefenseSuppliedTime != "" {
			t.Fatal("missing times must remain empty")
		}
		encoded, err := json.Marshal(n)
		if err != nil {
			t.Fatal(err)
		}
		var out map[string]interface{}
		if err := json.Unmarshal(encoded, &out); err != nil {
			t.Fatal(err)
		}
		for _, key := range []string{"disputeCancelledTime", "defenseSuppliedTime"} {
			if _, exists := out[key]; exists {
				t.Fatalf("absent time emitted: %s", encoded)
			}
		}
	}
}
