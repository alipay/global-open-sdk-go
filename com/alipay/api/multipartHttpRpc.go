package defaultAlipayClient

import (
	"bytes"
	"fmt"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/exception"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"strings"
)

func doMultipartRequest(
	requestUrl string,
	httpMethod string,
	headers map[string]string,
	requestBody string,
	filePartNames []string,
	fileName string,
	fileContentType string,
	fileContent []byte,
) (http.Header, []byte, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)
	if err := writeMultipartTextPart(writer, "body", requestBody); err != nil {
		return nil, nil, &exception.AlipayLibraryError{Message: "build multipart body is fail " + err.Error()}
	}
	for _, fieldName := range filePartNames {
		if err := writeMultipartFilePart(writer, fieldName, fileName, fileContentType, fileContent); err != nil {
			return nil, nil, &exception.AlipayLibraryError{Message: "build multipart file is fail " + err.Error()}
		}
	}
	if err := writer.Close(); err != nil {
		return nil, nil, &exception.AlipayLibraryError{Message: "close multipart writer is fail " + err.Error()}
	}

	req, err := http.NewRequest(httpMethod, requestUrl, bytes.NewReader(payload.Bytes()))
	if err != nil {
		return nil, nil, &exception.AlipayLibraryError{Message: "http.NewRequest is fail " + err.Error()}
	}
	for key, value := range headers {
		if !strings.EqualFold(key, "Content-Type") {
			req.Header.Set(key, value)
		}
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	transport := &http.Transport{
		DialContext:           (&net.Dialer{Timeout: connectTimeout}).DialContext,
		ResponseHeaderTimeout: readTimeout,
	}
	client := &http.Client{Timeout: totalTimeout, Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, &exception.AlipayLibraryError{Message: "client.Do is fail " + err.Error()}
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, &exception.AlipayLibraryError{Message: "read response is fail " + err.Error()}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, nil, &exception.AlipayLibraryError{Message: fmt.Sprintf("invalid http status %d, response body: %s", resp.StatusCode, responseBody)}
	}
	return resp.Header, responseBody, nil
}

func writeMultipartTextPart(writer *multipart.Writer, fieldName string, value string) error {
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"`, fieldName))
	header.Set("Content-Type", "application/json; charset=UTF-8")
	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}
	_, err = io.WriteString(part, value)
	return err
}

func writeMultipartFilePart(writer *multipart.Writer, fieldName string, fileName string, contentType string, content []byte) error {
	header := make(textproto.MIMEHeader)
	header.Set(
		"Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldName, sanitizeMultipartFileName(fileName)),
	)
	header.Set("Content-Type", contentType)
	header.Set("Content-Transfer-Encoding", "binary")
	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}
	_, err = part.Write(content)
	return err
}

func sanitizeMultipartFileName(fileName string) string {
	if separator := strings.LastIndexAny(fileName, `/\\`); separator >= 0 {
		fileName = fileName[separator+1:]
	}
	return strings.Map(func(value rune) rune {
		if value == '"' || value < 0x20 || value == 0x7f {
			return '_'
		}
		return value
	}, fileName)
}
