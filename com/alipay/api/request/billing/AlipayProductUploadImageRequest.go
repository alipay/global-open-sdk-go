package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	"io"
	"os"
	"path/filepath"
)

type AlipayProductUploadImageRequest struct {
	ProductId  string    `json:"productId,omitempty"`
	FileReader io.Reader `json:"-"`
	Filename   string    `json:"-"`
}

func NewAlipayProductUploadImageRequest() (*request.AlipayFileRequest, *AlipayProductUploadImageRequest) {
	alipayProductUploadImageRequest := &AlipayProductUploadImageRequest{}
	alipayRequest := request.NewAlipayFileRequest(alipayProductUploadImageRequest)
	return alipayRequest, alipayProductUploadImageRequest
}

func (alipayProductUploadImageRequest *AlipayProductUploadImageRequest) NewRequest() *request.AlipayFileRequest {
	return request.NewAlipayFileRequest(alipayProductUploadImageRequest)
}

// SetFile configures a local file without transferring ownership to the SDK.
func (alipayProductUploadImageRequest *AlipayProductUploadImageRequest) SetFile(file *os.File) {
	alipayProductUploadImageRequest.FileReader = file
	if file == nil {
		alipayProductUploadImageRequest.Filename = ""
		return
	}
	alipayProductUploadImageRequest.Filename = filepath.Base(file.Name())
}
