# Independent API Key client

This source adds `ApiKeyAlipayClient` to the existing SDK package. It does not change the RSA client or business models. It is not a statement of a released minimum version.

## Initialization and ordinary calls

The regional HTTPS gateway and complete Standard/Restricted API Key are the only required client settings. Read secrets from server-side configuration; never log the key or Authorization header. No ClientId decoding is performed. TEST uses `/ams/sandbox/api/`; PROD uses `/ams/api/`. The Request is not modified.

The following is a client integration fragment; business fields remain product-specific.

```go
import (
    "os"
    antom "github.com/alipay/global-open-sdk-go/com/alipay/api"
    pay "github.com/alipay/global-open-sdk-go/com/alipay/api/request/pay"
)

client, err := antom.NewApiKeyAlipayClient(os.Getenv("ANTOM_GATEWAY_URL"), os.Getenv("ANTOM_API_KEY"))
if err != nil { return err }
defer client.Close()
request, business := pay.NewAlipayPaymentSessionRequest()
// Populate business with the required payment-product fields before execution.
_ = business
response, err := client.Execute(request)
// Handle err and the existing typed business response as with RSA.
_ = response
```

## Configuration and lifecycle

Optional `WithApiKeyTimeout(30 * time.Second)`. The client reuses connections; call `Close()` to release idle connections. The normal environment proxy configuration is honored.

Ordinary requests validate TLS certificates and hostnames, do not follow redirects, and do not add automatic retries. Custom headers cannot override Bearer, RSA authentication fields or SDK transport headers.

## Migration and special calls

Existing RSA users keep their client, constructors and calls. Opting into API Key requires changing imports, initialization and any declarations tied to the concrete RSA client. Custom RSA subclasses do not automatically migrate. Business request models and result handling stay the same.

Use the language’s existing extra-header call convention on the new concrete client for `uploadEvent`, supplying only `X-Session-Id`. The existing Session HTTP/2 executor handles this route without Bearer or ordinary environment-path rewriting. This does not establish Billing/Meter API Key permissions.

File upload continues to require a separate RSA client. Notifications continue to use the existing RSA notification tool with independently configured matching ClientId and public key. API Key is not a notification signing key; HMAC is outside this change.

## Maintenance

The new Client, API Key authentication helper and any API Key transport files are hand-maintained. Existing generation deploys business models/requests/responses into their own directories. Keep these files tracked before running generator cleanup; `git clean` intentionally removes untracked files.
