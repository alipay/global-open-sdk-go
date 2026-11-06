package defaultAlipayClient

const SDKVersion = "1.5.3"

func sdkUserAgent() string {
	return "global-open-sdk-go/" + SDKVersion
}
