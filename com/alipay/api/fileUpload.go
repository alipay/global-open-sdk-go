package defaultAlipayClient

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/exception"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	requestBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/request/billing"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/tools"
	"io"
	"mime"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const productImageMaxFileSize int64 = 2 * 1024 * 1024

type fileUploadOperation struct {
	path          string
	httpMethod    string
	maxFileSize   int64
	filePartNames []string
	prepare       func(any) (io.Reader, string, map[string]string, error)
	newResponse   func() any
}

var productImageUploadOperation = fileUploadOperation{
	path:          "/ams/api/v1/billing/product/uploadImage",
	httpMethod:    model.MethodPost,
	maxFileSize:   productImageMaxFileSize,
	filePartNames: []string{"file", "imageFile"},
	prepare:       prepareProductImageUpload,
	newResponse: func() any {
		return &responseBilling.AlipayProductUploadImageResponse{}
	},
}

func executeFileUpload(client *DefaultAlipayClient, alipayRequest *request.AlipayFileRequest) (any, error) {
	if alipayRequest == nil {
		return nil, libraryError("alipayFileRequest can't be nil")
	}
	operation, err := resolveFileUploadOperation(alipayRequest.Param)
	if err != nil {
		return nil, err
	}
	reader, filename, businessBody, err := operation.prepare(alipayRequest.Param)
	if err != nil {
		return nil, err
	}
	fileContent, err := readBoundedFile(reader, operation.maxFileSize)
	if err != nil {
		return nil, err
	}
	digest := sha256.Sum256(fileContent)
	businessBody["fileSha256"] = fmt.Sprintf("%x", digest)
	requestBodyBytes, err := json.Marshal(businessBody)
	if err != nil {
		return nil, libraryError("json.Marshal file request body failed: " + err.Error())
	}
	requestBody := string(requestBodyBytes)

	clientID := alipayRequest.ClientId
	if strings.TrimSpace(clientID) == "" {
		clientID = client.ClientId
	}
	if strings.TrimSpace(clientID) == "" {
		return nil, libraryError("clientId can't be empty")
	}
	uploadGateway, err := resolveUploadGateway(client.GatewayUrl, client.uploadGatewayUrl)
	if err != nil {
		return nil, err
	}
	path := resolveFileUploadPath(operation.path, strings.HasPrefix(clientID, "SANDBOX_"))
	requestTime := strconv.FormatInt(time.Now().UnixMilli(), 10)
	signature, err := tools.GenSign(
		operation.httpMethod,
		path,
		clientID,
		requestTime,
		requestBody,
		client.MerchantPrivateKey,
	)
	if err != nil {
		return nil, err
	}
	headers := buildBaseHeader(
		requestTime,
		clientID,
		alipayRequest.KeyVersion,
		signature,
		client.AgentToken,
	)
	contentType := mime.TypeByExtension(filepath.Ext(filename))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	responseHeaders, responseBody, err := doMultipartRequest(
		uploadGateway+path,
		operation.httpMethod,
		headers,
		requestBody,
		operation.filePartNames,
		filename,
		contentType,
		fileContent,
	)
	if err != nil {
		return nil, err
	}
	response := operation.newResponse()
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, libraryError("json.Unmarshal file response failed: " + err.Error())
	}
	var resultEnvelope struct {
		Result *model.Result `json:"result"`
	}
	if err := json.Unmarshal(responseBody, &resultEnvelope); err != nil || resultEnvelope.Result == nil {
		return nil, libraryError("file response result field is missing")
	}

	responseSignature := responseHeaders.Get("Signature")
	responseTime := responseHeaders.Get("Response-Time")
	if responseSignature == "" && responseTime == "" {
		if resultEnvelope.Result.ResultStatus == model.ResultStatusType_F {
			return response, nil
		}
		return nil, libraryError("unsigned file response is not a failure response")
	}
	if responseSignature == "" || responseTime == "" {
		return nil, libraryError("file response must contain both Signature and Response-Time")
	}
	verified, err := checkRspSign(
		operation.httpMethod,
		path,
		clientID,
		responseTime,
		string(responseBody),
		responseSignature,
		client.AlipayPublicKey,
	)
	if err != nil {
		return nil, err
	}
	if !verified {
		return nil, libraryError("response signature verification failed")
	}
	return response, nil
}

func resolveFileUploadOperation(param any) (*fileUploadOperation, error) {
	switch param.(type) {
	case *requestBilling.AlipayProductUploadImageRequest:
		return &productImageUploadOperation, nil
	default:
		return nil, libraryError("only SDK-provided file upload requests are supported")
	}
}

func prepareProductImageUpload(param any) (io.Reader, string, map[string]string, error) {
	productRequest, ok := param.(*requestBilling.AlipayProductUploadImageRequest)
	if !ok || productRequest == nil {
		return nil, "", nil, libraryError("invalid product image upload request")
	}
	if strings.TrimSpace(productRequest.ProductId) == "" {
		return nil, "", nil, libraryError("ProductId can't be empty")
	}
	if len(productRequest.ProductId) > 64 {
		return nil, "", nil, libraryError("ProductId length can't exceed 64 characters")
	}
	if productRequest.FileReader == nil {
		return nil, "", nil, libraryError("FileReader can't be nil")
	}
	filename := sanitizeMultipartFileName(strings.TrimSpace(productRequest.Filename))
	if filename == "" {
		return nil, "", nil, libraryError("Filename can't be empty")
	}
	return productRequest.FileReader, filename, map[string]string{
		"productId": productRequest.ProductId,
	}, nil
}

func readBoundedFile(reader io.Reader, maxFileSize int64) ([]byte, error) {
	var seeker io.Seeker
	var originalPosition int64
	if candidate, ok := reader.(io.Seeker); ok {
		position, err := candidate.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, libraryError("read file position failed: " + err.Error())
		}
		seeker = candidate
		originalPosition = position
	}
	content, readErr := io.ReadAll(io.LimitReader(reader, maxFileSize+1))
	if seeker != nil {
		if _, err := seeker.Seek(originalPosition, io.SeekStart); err != nil {
			return nil, libraryError("restore file position failed: " + err.Error())
		}
	}
	if readErr != nil {
		return nil, libraryError("read file failed: " + readErr.Error())
	}
	if len(content) == 0 {
		return nil, libraryError("file can't be empty")
	}
	if int64(len(content)) > maxFileSize {
		return nil, libraryError(fmt.Sprintf("file size can't exceed %d bytes", maxFileSize))
	}
	return content, nil
}

func resolveFileUploadPath(path string, sandbox bool) string {
	if sandbox && !shouldUseProductionPathInSandbox(path) {
		return strings.Replace(path, "/ams/api", "/ams/sandbox/api", 1)
	}
	return path
}

func libraryError(message string) error {
	return &exception.AlipayLibraryError{Message: message}
}
