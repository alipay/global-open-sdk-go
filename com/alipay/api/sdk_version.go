package defaultAlipayClient

const SDKVersion = "1.4.1"

func sdkUserAgent() string {
	return "global-open-sdk-go/" + SDKVersion
}
