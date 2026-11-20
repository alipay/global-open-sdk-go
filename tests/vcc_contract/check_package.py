"""Go source-package inventory and a real consumer outside the SDK checkout."""
import hashlib
import json
import os
from pathlib import Path
import subprocess
import sys


def run(argv, cwd):
    p = subprocess.run(argv, cwd=cwd, text=True, capture_output=True)
    if p.returncode:
        raise RuntimeError(str(argv) + '\n' + p.stdout[-3000:] + p.stderr[-3000:])
    return p.stdout


evidence = json.loads(Path(os.environ['IL_SDK_BUILD_EVIDENCE']).read_text())
assert evidence['native_status'] == 'passed'
root = Path(evidence['workspace']).resolve()
module = 'github.com/alipay/global-open-sdk-go'
mode = sys.argv[1]
if mode == 'inventory':
    assert Path.cwd().resolve() == root
    assert ['go', 'test', './...'] in evidence['commands']
    packages = set(run(['go', 'list', './...'], root).splitlines())
    for stem in ['auth_demo', 'customs_demo', 'dispute', 'marketplace_demo', 'pay_demo', 'subscription_demo', 'vaulting_demo']:
        before = 'com/alipay/example/' + stem + '.go'
        after = root / 'com/alipay/example' / stem / (stem + '.go')
        baseline = subprocess.check_output(['git', 'show', evidence['baseline_commit'] + ':' + before], cwd=root)
        assert hashlib.sha256(after.read_bytes()).digest() == hashlib.sha256(baseline).digest()
        assert module + '/com/alipay/example/' + stem in packages
    assert module + '/com/alipay/api/model' in packages
    print('All seven unchanged demo bodies are present in separate compiled packages; native full-repository build reused.')
elif mode == 'consumer':
    consumer = Path.cwd().resolve()
    assert consumer != root and root not in consumer.parents
    (consumer / 'go.mod').write_text('module vcc-consumer\n\ngo 1.22.5\n\nrequire ' + module + ' v0.0.0\nreplace ' + module + ' => ' + json.dumps(str(root)) + '\n')
    (consumer / 'consumer_test.go').write_text('''package consumer
import (
 "encoding/json"
 "testing"
 "reflect"
 m "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 a "github.com/alipay/global-open-sdk-go/com/alipay/api/request/aba"
 r "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)
func TestInstalledSourceConsumer(t *testing.T) {
 currencies:=[]string{"HKD","USD"}
 request,body:=a.NewAlipayApplyCardRequest()
 body.AuthorizationControl=&m.ApplyCardAuthorizationControl{PaymentPreference:&m.ApplyCardPaymentPreference{CurrencyOrder:&currencies}}
 data,err:=json.Marshal(body);if err!=nil {t.Fatal(err)}
 var wire map[string]any;if err=json.Unmarshal(data,&wire);err!=nil {t.Fatal(err)}
 actual:=wire["authorizationControl"].(map[string]any)["paymentPreference"].(map[string]any)["currencyOrder"]
 if !reflect.DeepEqual(actual,[]any{"HKD","USD"}) {t.Fatal(string(data))}
 if request.HttpMethod!="POST" {t.Fatal(request.HttpMethod)}
 var reply r.AlipayInquireCardSensitiveInfoResponse
 if err=json.Unmarshal([]byte(`{"cardDetail":{"authorizationControl":{"paymentPreference":{"currencyOrder":["HKD","USD"]}}}}`),&reply);err!=nil {t.Fatal(err)}
 if !reflect.DeepEqual(*reply.CardDetail.AuthorizationControl.PaymentPreference.CurrencyOrder,currencies) {t.Fatal("response model binding")}
}
''')
    run(['go', 'mod', 'tidy'], consumer)
    print(run(['go', 'test', './...'], consumer))
else:
    raise AssertionError(mode)
