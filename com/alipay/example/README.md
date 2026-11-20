# Standalone examples

Each demo keeps its own `main` function in a separate package so all examples remain runnable and `go test ./...` and `go vet ./...` can cover the complete repository.

Run a demo from the repository root, for example:

```sh
go run ./com/alipay/example/auth_demo
go run ./com/alipay/example/pay_demo
```

The other standalone demos use the same directory convention. Configure the placeholders in the selected demo before a live request. Contract tests do not execute these live-request entry points.
