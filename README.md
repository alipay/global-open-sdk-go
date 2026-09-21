# Antom SDK for Go

Latest release: **1.4.1**

## Installation

```sh
go get github.com/alipay/global-open-sdk-go@latest
```

Requires Go 1.22.5+. Use a supported Go release with current security fixes.

## Quick start

- **API Key:** follow the [setup guide](docs/api-key-client.md) and run the [sandbox example](com/alipay/example/api_key_payment_session/main.go).
- **RSA:** start with the [payment example](com/alipay/example/pay_demo.go).
- Browse [more examples](com/alipay/example) and the [API documentation](https://global.alipay.com/docs/).

API Key and RSA clients share request/response models. File uploads and notification
verification still require RSA credentials.

## Upgrade notes

Billing integrations: `availableAmount` now uses `Amount`; the `AvailableAmount`
model has been removed.

## Meter event upload

`meter/uploadEvent` requires HTTP/2 and `X-Session-Id`. See the
[usage and requirements](docs/meter-event-upload.md).

## Support

For integration questions, contact overseas_support@service.alibaba.com.
