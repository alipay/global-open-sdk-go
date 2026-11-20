package main

import (
	"encoding/json"
	"fmt"
	m "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	a "github.com/alipay/global-open-sdk-go/com/alipay/api/request/aba"
	r "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
	"os"
	"reflect"
	"strings"
)

type S = map[string]any

func must(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}
func deref(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}
func schema(t reflect.Type) S {
	t = deref(t)
	switch t.Kind() {
	case reflect.String:
		return S{"type": "string"}
	case reflect.Bool:
		return S{"type": "boolean"}
	case reflect.Int, reflect.Int32, reflect.Int64:
		return S{"type": "integer"}
	case reflect.Float32, reflect.Float64:
		return S{"type": "number"}
	case reflect.Slice, reflect.Array:
		return S{"type": "array", "items": schema(t.Elem())}
	case reflect.Map:
		return S{"type": "object", "properties": S{}, "additionalProperties": schema(t.Elem())}
	default:
		return S{"$ref": "#/components/schemas/" + t.Name()}
	}
}
func value(t reflect.Type) reflect.Value {
	if t.Kind() == reflect.Pointer {
		v := reflect.New(t.Elem())
		v.Elem().Set(value(t.Elem()))
		return v
	}
	v := reflect.New(t).Elem()
	switch t.Kind() {
	case reflect.String:
		v.SetString("probe")
	case reflect.Bool:
		v.SetBool(false)
	case reflect.Int, reflect.Int32, reflect.Int64:
		v.SetInt(1)
	case reflect.Slice:
		v = reflect.MakeSlice(t, 2, 2)
		v.Index(0).SetString("HKD")
		v.Index(1).SetString("USD")
	case reflect.Map:
		v = reflect.MakeMap(t)
		v.SetMapIndex(reflect.ValueOf("probe"), reflect.ValueOf("value"))
	}
	return v
}
func wire(obj any) any {
	b, e := json.Marshal(obj)
	if e != nil {
		panic(e)
	}
	var out any
	e = json.Unmarshal(b, &out)
	if e != nil {
		panic(e)
	}
	return out
}
func name(v any) string { return deref(reflect.TypeOf(v)).Name() }
func main() {
	var req struct {
		Models []string `json:"models"`
	}
	if e := json.NewDecoder(os.Stdin).Decode(&req); e != nil {
		panic(e)
	}
	catalog := map[string]reflect.Type{}
	for _, obj := range []any{a.AlipayApplyCardRequest{}, r.AlipayInquireCardDetailResponse{}, m.ApplyCardAuthorizationControl{}, m.ApplyCardPaymentPreference{}, m.CardDetail{}, m.CardInquiryAuthorizationControl{}, m.CardPaymentPreference{}} {
		catalog[name(obj)] = reflect.TypeOf(obj)
	}
	models := S{}
	for _, n := range req.Models {
		t, ok := catalog[n]
		must(ok, "missing actual type "+n)
		obj := reflect.New(t)
		props := S{}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Anonymous {
				continue
			}
			key := strings.Split(f.Tag.Get("json"), ",")[0]
			if key == "-" {
				continue
			}
			if key == "" {
				key = f.Name
			}
			props[key] = schema(f.Type)
			obj.Elem().Field(i).Set(value(f.Type))
		}
		body := wire(obj.Interface())
		must(len(body.(map[string]any)) == len(props), "all actual keys "+n)
		data, _ := json.Marshal(obj.Interface())
		round := reflect.New(t)
		must(json.Unmarshal(data, round.Interface()) == nil, "parse "+n)
		must(reflect.DeepEqual(wire(round.Interface()), body), "roundtrip "+n)
		models[n] = S{"type": "object", "properties": props}
	}
	empty := []string{}
	custom := []string{"HKD", "USD"}
	must(len(wire(m.ApplyCardPaymentPreference{}).(map[string]any)) == 0, "absent child")
	must(len(wire(m.ApplyCardPaymentPreference{CurrencyOrder: &empty}).(map[string]any)["currencyOrder"].([]any)) == 0, "explicit empty array")
	must(reflect.DeepEqual(wire(m.ApplyCardPaymentPreference{CurrencyOrder: &custom}).(map[string]any)["currencyOrder"], []any{"HKD", "USD"}), "ordered array")
	must(reflect.DeepEqual(wire(m.ApplyCardAuthorizationControl{PaymentPreference: &m.ApplyCardPaymentPreference{}}).(map[string]any)["paymentPreference"], map[string]any{}), "inheritance object")
	var ordinary r.AlipayInquireCardDetailResponse
	var sensitive r.AlipayInquireCardSensitiveInfoResponse
	must(json.Unmarshal([]byte(`{"authorizationControl":{"paymentPreference":{"currencyOrder":["HKD","USD"]}}}`), &ordinary) == nil, "ordinary parse")
	must(json.Unmarshal([]byte(`{"cardDetail":{"authorizationControl":{"paymentPreference":{"currencyOrder":["HKD","USD"]}}}}`), &sensitive) == nil, "sensitive parse")
	must(reflect.DeepEqual(*ordinary.AuthorizationControl.PaymentPreference.CurrencyOrder, custom), "ordinary binding")
	must(reflect.DeepEqual(*sensitive.CardDetail.AuthorizationControl.PaymentPreference.CurrencyOrder, custom), "sensitive binding")
	q1, p1 := a.NewAlipayApplyCardRequest()
	q2, p2 := a.NewAlipayInquireCardDetailRequest()
	q3, p3 := a.NewAlipayInquireCardSensitiveInfoRequest()
	bindings := []S{{"path": q1.Path, "method": q1.HttpMethod, "request": name(p1), "response": name(q1.AlipayResponse)}, {"path": q2.Path, "method": q2.HttpMethod, "request": name(p2), "response": name(q2.AlipayResponse)}, {"path": q3.Path, "method": q3.HttpMethod, "request": name(p3), "response": name(q3.AlipayResponse)}}
	out := S{"models": models, "bindings": bindings, "wire_checks": len(models), "evidence": []string{"Native Go reflect struct/tag/type observation, encoding/json roundtrip, pointer-to-array empty vs absent, nested inquiry parsing and real request constructors"}}
	data, e := json.Marshal(out)
	if e != nil {
		panic(e)
	}
	fmt.Print(string(data))
}
