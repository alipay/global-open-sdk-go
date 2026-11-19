import json,pathlib,re,subprocess,sys,tempfile
root=pathlib.Path.cwd();req=json.load(sys.stdin);wanted=set(req['models']);operations=set()
for t in req['operations']:
 op=t.split(' ',1)[1].split('/requestBody',1)[0].split('/responses',1)[0].rsplit('/',1)[1];operations.add(op);wanted.add('Alipay'+op[0].upper()+op[1:]+'Request');wanted.add('Alipay'+op[0].upper()+op[1:]+'Response')
module=re.search(r'^module\s+(\S+)',(root/'go.mod').read_text(),re.M).group(1);imports={};names={};sources={}
for p in (root/'com/alipay/api').rglob('*.go'):
 for name in re.findall(r'\btype\s+(\w+)\s+struct\s*\{',p.read_text()):
  if name in wanted:
   package=p.parent.relative_to(root).as_posix();alias=imports.setdefault(package,'p'+str(len(imports)));names[name]=alias+'.'+name;sources[name]=p
assert set(names)==wanted,(set(names),wanted)
code='package main\nimport ("encoding/json";"fmt";"os";"reflect";"strings";'
code+=';'.join(alias+' "'+module+'/'+path+'"' for path,alias in imports.items())+')\n'
code+=r'''
func object(kv ...interface{})map[string]interface{}{m:=map[string]interface{}{};for i:=0;i<len(kv);i+=2{m[kv[i].(string)]=kv[i+1]};return m}
func shape(t reflect.Type)interface{}{
 for t.Kind()==reflect.Ptr{t=t.Elem()}
 if strings.HasPrefix(t.Name(),"Nullable"){return object("$ref","#/components/schemas/"+strings.TrimPrefix(t.Name(),"Nullable"))}
 switch t.Kind(){case reflect.String:return object("type","string");case reflect.Bool:return object("type","boolean");case reflect.Int,reflect.Int32,reflect.Int64,reflect.Uint,reflect.Uint64:return object("type","integer");case reflect.Float32,reflect.Float64:return object("type","number");case reflect.Slice,reflect.Array:return object("type","array","items",shape(t.Elem()));case reflect.Map:return object("type","object","additionalProperties",shape(t.Elem()),"properties",object());case reflect.Interface:return object();default:return object("$ref","#/components/schemas/"+t.Name())}
}
func sample(t reflect.Type,depth int)interface{}{
 if depth>8{return nil};for t.Kind()==reflect.Ptr{t=t.Elem()}
 switch t.Kind(){case reflect.String:return "USD";case reflect.Bool:return true;case reflect.Int,reflect.Int32,reflect.Int64,reflect.Uint,reflect.Uint64:return 1;case reflect.Float32,reflect.Float64:return 1.5;case reflect.Slice,reflect.Array:if t.Elem().Kind()==reflect.String{return []string{"HKD","USD"}};return []interface{}{sample(t.Elem(),depth+1)};case reflect.Map:return object("key",sample(t.Elem(),depth+1));case reflect.Struct:m:=map[string]interface{}{};for i:=0;i<t.NumField();i++{f:=t.Field(i);if f.PkgPath!=""{continue};key:=strings.Split(f.Tag.Get("json"),",")[0];if key=="-"||key==""{continue};m[key]=sample(f.Type,depth+1)};return m;default:return "value"}
}
func observe(t reflect.Type)interface{}{
 props:=map[string]interface{}{};for i:=0;i<t.NumField();i++{f:=t.Field(i);key:=strings.Split(f.Tag.Get("json"),",")[0];if f.PkgPath!=""||key=="-"||key==""{continue};props[key]=shape(f.Type)}
 input,e:=json.Marshal(sample(t,0));if e!=nil{panic(e)};v:=reflect.New(t);if e=json.Unmarshal(input,v.Interface());e!=nil{panic(e)};encoded,e:=json.Marshal(v.Interface());if e!=nil{panic(e)}
 var got map[string]interface{};json.Unmarshal(encoded,&got);for key:=range props{if _,ok:=got[key];!ok{panic("serializer dropped "+t.Name()+"."+key)}}
 parsed:=reflect.New(t);if e=json.Unmarshal(encoded,parsed.Interface());e!=nil{panic(e)};again,_:=json.Marshal(parsed.Interface());if string(again)!=string(encoded){panic("roundtrip "+t.Name())}
 if strings.HasSuffix(t.Name(),"PaymentPreference") {
 for _,raw:=range []string{`{"currencyOrder":["HKD","USD"]}`,`{"currencyOrder":["USD","HKD"]}`,`{"currencyOrder":[]}`} {
  policy:=reflect.New(t);if e=json.Unmarshal([]byte(raw),policy.Interface());e!=nil{panic(e)};encoded,_:=json.Marshal(policy.Interface());
  var before,after interface{};json.Unmarshal([]byte(raw),&before);json.Unmarshal(encoded,&after);if !reflect.DeepEqual(before,after){panic("ordered/empty currencyOrder lost: "+string(encoded))}
 }
 }
 return object("type","object","properties",props)
}
func main(){models:=map[string]interface{}{};bindings:=map[string]interface{}{}
'''
for name in req['models']:code+='models['+json.dumps(name)+']=observe(reflect.TypeOf('+names[name]+'{}))\n'
for op in sorted(operations):
 name='Alipay'+op[0].upper()+op[1:]+'Request';txt=sources[name].read_text();paths=re.findall(r'"(/ams/api/[^"\n]+)"',txt);assert len(set(paths))==1
 response=name[:-7]+'Response';assert re.search(r'\b'+re.escape(response)+r'\b',txt)
 code+='bindings['+json.dumps(op)+']=object("path",'+json.dumps(paths[0])+',"request",reflect.TypeOf('+names[name]+'{}).Name(),"response",reflect.TypeOf('+names[response]+'{}).Name())\n'
code+='json.NewEncoder(os.Stdout).Encode(object("models",models,"bindings",bindings,"evidence",[]string{"reflect.Type field/tag/generic shape; actual encoding/json unmarshal, marshal and second typed roundtrip for populated models"}));fmt.Sprint()\n}\n'
# Independent main programs cannot share one Go package. Keep their explicit
# `go run file.go` entry points and compile/vet every program, not just the SDK.
examples=sorted((root/'com/alipay/example').glob('*.go'))
with tempfile.TemporaryDirectory(prefix='vcc-go-examples-') as temp:
 for source in examples:
  if not re.search(r'func\s+main\s*\(',source.read_text()):continue
  for command in [['go','build','-o',str(pathlib.Path(temp)/(source.stem+'-example')),str(source)],['go','vet',str(source)]]:
   checked=subprocess.run(command,cwd=root,capture_output=True,text=True)
   assert checked.returncode==0,checked.stderr[-4000:]
with tempfile.TemporaryDirectory(prefix='vcc-go-observer-') as tmp:
 path=pathlib.Path(tmp)/'main.go';path.write_text(code);p=subprocess.run(['go','run',str(path)],cwd=root,capture_output=True,text=True,check=False);assert p.returncode==0,p.stderr[-4000:]+p.stdout[-1000:];actual=json.loads(p.stdout);actual['evidence'].append(str(sum(bool(re.search(r'func\s+main\s*\(',x.read_text())) for x in examples))+' standalone examples compiled and vetted individually');print(json.dumps(actual))
