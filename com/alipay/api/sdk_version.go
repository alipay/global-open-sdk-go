package defaultAlipayClient

const SDKVersion = "1.3.8"

func sdkUserAgent() string {
	return "global-open-sdk-go/" + SDKVersion
}
