# Standalone examples

Each executable demo keeps its own `main` function in a separate package. The root `example` package provides reusable helpers such as `UploadMeterEvents` without an executable entry point.

Validate the complete repository with `go build ./...`, `go test ./...`, and `go vet ./...`.

Run a demo from the repository root, for example:

```sh
go run ./com/alipay/example/auth_demo
go run ./com/alipay/example/pay_demo
```

The other standalone demos use the same directory convention. Configure the placeholders in the selected demo before a live request. Contract tests do not execute these live-request entry points.
