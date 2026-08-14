package defaultAlipayClient

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alipay/global-open-sdk-go/com/alipay/api/exception"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
)

const sessionHeader = "X-Session-Id"

func executeSessionHTTP2(gatewayURL string, alipayRequest *request.AlipayRequest, extraHeaders map[string]string) (any, error) {
	sessionID, err := validateAndGetSessionID(extraHeaders)
	if err != nil {
		return nil, err
	}

	requestBody, err := json.Marshal(alipayRequest.Param)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: "json.Marshal failed: " + err.Error()}
	}
	responseBody, err := postHTTP2JSON(gatewayURL, alipayRequest.Path, sessionID, requestBody)
	if err != nil {
		return nil, err
	}
	if alipayRequest.AlipayResponse == nil {
		return nil, &exception.AlipayLibraryError{Message: "AlipayResponse cannot be nil"}
	}
	if err := json.Unmarshal(responseBody, alipayRequest.AlipayResponse); err != nil {
		return nil, &exception.AlipayLibraryError{Message: "json.Unmarshal failed: " + err.Error()}
	}

	var envelope struct {
		Result json.RawMessage `json:"result"`
	}
	if err := json.Unmarshal(responseBody, &envelope); err != nil || len(envelope.Result) == 0 || string(envelope.Result) == "null" {
		return nil, &exception.AlipayLibraryError{Message: "response data error: result field is null"}
	}
	return alipayRequest.AlipayResponse, nil
}

func validateAndGetSessionID(extraHeaders map[string]string) (string, error) {
	sessionID := ""
	sessionHeaderFound := false
	for name, value := range extraHeaders {
		if !strings.EqualFold(name, sessionHeader) {
			return "", &exception.AlipayLibraryError{Message: fmt.Sprintf("only X-Session-Id is supported for this API; unsupported header: %s", name)}
		}
		if sessionHeaderFound {
			return "", &exception.AlipayLibraryError{Message: "X-Session-Id must be provided only once"}
		}
		sessionHeaderFound = true
		sessionID = value
	}
	if strings.TrimSpace(sessionID) == "" {
		return "", &exception.AlipayLibraryError{Message: "X-Session-Id cannot be empty"}
	}
	if strings.ContainsAny(sessionID, "\r\n") {
		return "", &exception.AlipayLibraryError{Message: "X-Session-Id cannot contain CR or LF characters"}
	}
	return sessionID, nil
}
