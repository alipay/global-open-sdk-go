# Meter event upload

`meter/createSession` uses the regular signed AMS transport. Call
`meter/uploadEvent` with `ExecuteWithHeaders` and the returned `X-Session-Id`.
The SDK uses the gateway URL configured on the client without sandbox path
rewriting, request signing, response signature verification, or automatic
retries. This API requires HTTP/2 and supports the Go version declared in
`go.mod`. For production use, build applications with a currently supported
Go release containing the latest security fixes.

See the [Meter example](../com/alipay/example/meter_upload_event_demo.go) for a complete request.
