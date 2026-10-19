package account_acceptance

import (
	"encoding/json"
	n "github.com/alipay/global-open-sdk-go/com/alipay/api/request/notify"
	p "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
	"os"
	"reflect"
	"testing"
)

func TestAccountOutputs(t *testing.T) {
	data, e := os.ReadFile("testdata/account-cases.json")
	if e != nil {
		t.Fatal(e)
	}
	var cases []map[string]json.RawMessage
	if e = json.Unmarshal(data, &cases); e != nil {
		t.Fatal(e)
	}
	for _, c := range cases {
		body, _ := json.Marshal(c)
		for _, o := range []interface{}{&p.AlipayPayResponse{}, &p.AlipayPayQueryResponse{}, &p.AlipayCaptureResponse{}, &n.AlipayPayResultNotify{}, &n.AlipayCaptureResultNotify{}} {
			if e = json.Unmarshal(body, o); e != nil {
				t.Fatal(e)
			}
			a := reflect.ValueOf(o).Elem().FieldByName("AcquirerInfo")
			want, ok := c["acquirerInfo"]
			if !ok {
				if !a.IsNil() {
					t.Fatal("absent")
				}
				continue
			}
			enc, e := json.Marshal(a.Interface())
			if e != nil {
				t.Fatal(e)
			}
			var actual, expected interface{}
			json.Unmarshal(enc, &actual)
			json.Unmarshal(want, &expected)
			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("%T mismatch %s", o, enc)
			}
		}
	}
}
