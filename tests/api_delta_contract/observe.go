package main

import (
	"encoding/json"
	"fmt"
	pkg1 "github.com/alipay/global-open-sdk-go/com/alipay/api/request/billing"
	pkg0 "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"

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
	if t.PkgPath() != "" && t.Kind() != reflect.Struct {
		return S{"$ref": "#/components/schemas/" + t.Name()}
	}
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
		v.SetBool(true)
	case reflect.Int, reflect.Int32, reflect.Int64:
		v.SetInt(1)
	case reflect.Slice:
		v = reflect.MakeSlice(t, 2, 2)
		v.Index(0).Set(value(t.Elem()))
		v.Index(1).Set(value(t.Elem()))
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
		Models        []string `json:"models"`
		OperationKeys []string `json:"operation_keys"`
	}
	if e := json.NewDecoder(os.Stdin).Decode(&req); e != nil {
		panic(e)
	}
	catalog := map[string]reflect.Type{}
	for _, obj := range []any{pkg0.AlipayBillingSubscriptionUpdateResponse{}} {
		catalog[name(obj)] = reflect.TypeOf(obj)
	}

	models := S{}
	for _, target := range req.Models {
		parts := strings.Split(target, "/")
		n := parts[len(parts)-1]
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
		models[target] = S{"type": "object", "properties": props}
	}
	evidence := []string{"Actual Go reflected fields/tags/types and encoding/json roundtrips"}
	if t, ok := catalog["ConsultBuyer"]; ok {
		for _, flag := range []string{"true", "false"} {
			o := reflect.New(t)
			f := o.Elem().FieldByName("BuyerAccessGroupFlag")
			v := reflect.ValueOf(flag)
			if f.Kind() == reflect.Pointer {
				ptr := reflect.New(f.Type().Elem())
				ptr.Elem().Set(v.Convert(f.Type().Elem()))
				f.Set(ptr)
			} else {
				f.Set(v.Convert(f.Type()))
			}
			must(wire(o.Interface()).(map[string]any)["buyerAccessGroupFlag"] == flag, "buyer flag must remain a JSON string")
		}
	}
	for _, t := range catalog {
		if _, ok := t.FieldByName("CreditNoteIds"); !ok {
			continue
		}
		for _, fixture := range []string{`{}`, `{"creditNoteIds":null}`, `{"creditNoteIds":[]}`, `{"creditNoteId":"cn_a","creditNoteIds":["cn_a"],"creditNoteAmount":"3000","creditNoteCurrency":"USD"}`, `{"creditNoteIds":["cn_a","cn_z"],"creditNoteAmount":"8000","creditNoteCurrency":"USD"}`} {
			o := reflect.New(t)
			must(json.Unmarshal([]byte(fixture), o.Interface()) == nil, "subscription parser")
			actual := wire(o.Interface()).(map[string]any)
			var input map[string]any
			json.Unmarshal([]byte(fixture), &input)
			for k, v := range input {
				if v != nil {
					if a, ok := v.([]any); ok && len(a) == 0 {
						continue
					}
					must(reflect.DeepEqual(actual[k], v), "subscription field "+k)
				}
			}
		}
	}
	bindings := []S{}
	{
		q, p := pkg1.NewAlipayBillingSubscriptionUpdateRequest()
		bindings = append(bindings, S{"path": q.Path, "method": q.HttpMethod, "request": name(p), "response": name(q.AlipayResponse)})
	}
	out := S{"models": models, "bindings": bindings, "wire_checks": len(models), "evidence": evidence}
	data, e := json.Marshal(out)
	if e != nil {
		panic(e)
	}
	fmt.Print(string(data))
}
