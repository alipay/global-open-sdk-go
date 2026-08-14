package main

import (
	"fmt"
	"time"

	defaultAlipayClient "github.com/alipay/global-open-sdk-go/com/alipay/api"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	billingRequest "github.com/alipay/global-open-sdk-go/com/alipay/api/request/billing"
	billingResponse "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

// UploadMeterEvents uploads events with the session ID returned by
// meter/createSession. Applications calling this API must be rebuilt with Go
// 1.25.13+ in the Go 1.25 series, Go 1.26.6+ in the Go 1.26 series, or a later
// stable Go release. Other SDK APIs can continue to use the Go version declared
// in go.mod.
func UploadMeterEvents(client *defaultAlipayClient.DefaultAlipayClient, sessionID string) {
	alipayRequest, uploadRequest := billingRequest.NewAlipayMeterUploadEventRequest()
	uploadRequest.Meters = []*model.MeterEventBatch{
		{
			EventName: "api_requests",
			Events: []*model.Event{
				{
					IdempotencyKey: "meter-event-001",
					EventTimestamp: time.Now().UnixMilli(),
					Payload: &model.EventPayload{
						CustomerId: "customer-001",
						Value:      "1",
					},
				},
			},
		},
	}

	result, err := client.ExecuteWithHeaders(alipayRequest, map[string]string{
		"X-Session-Id": sessionID,
	})
	if err != nil {
		panic(err)
	}

	response := result.(*billingResponse.AlipayMeterUploadEventResponse)
	fmt.Printf("result=%+v retryAfter=%d errors=%+v\n", response.Result, response.RetryAfter, response.Errors)
}
