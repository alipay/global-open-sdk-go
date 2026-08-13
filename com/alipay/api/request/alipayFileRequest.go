package request

// AlipayFileRequest is the envelope accepted by DefaultAlipayClient.UploadFile.
// It deliberately carries no path or HTTP method; those are fixed by the SDK operation registry.
type AlipayFileRequest struct {
	Param      any
	ClientId   string
	KeyVersion string
}

func NewAlipayFileRequest(param any) *AlipayFileRequest {
	return &AlipayFileRequest{
		Param:      param,
		KeyVersion: "1",
	}
}
