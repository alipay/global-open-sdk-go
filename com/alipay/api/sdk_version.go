package defaultAlipayClient

const SDKVersion = "1.3.4"

func sdkUserAgent() string {
	return "global-open-sdk-go/" + SDKVersion
}
