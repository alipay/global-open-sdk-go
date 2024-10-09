package defaultAlipayClient

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/exception"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const readTimeout = 15 * time.Second
const connectTimeout = 15 * time.Second
const totalTimeout = 30 * time.Second

type DefaultAlipayClient struct {
	GatewayUrl         string
	ClientId           string
	MerchantPrivateKey string
	AlipayPublicKey    string
	IsSandboxMode      bool
}

/**
 * 构造方法
 * @param gatewayUrl 网关地址
 * @param clientId 应用ID
 * @param merchantPrivateKey 商户私钥
 * @param alipayPublicKey 支付宝公钥
 */
func NewDefaultAlipayClient(gatewayUrl string, clientId string, merchantPrivateKey string, alipayPublicKey string) *DefaultAlipayClient {
	isSandboxMode := false
	if strings.HasPrefix(clientId, "SANDBOX_") {
		isSandboxMode = true
	}

	return &DefaultAlipayClient{
		GatewayUrl:         gatewayUrl,
		ClientId:           clientId,
		MerchantPrivateKey: merchantPrivateKey,
		AlipayPublicKey:    alipayPublicKey,
		IsSandboxMode:      isSandboxMode,
	}
}

// @Description 支持任意方式的http请求
// @Param url 请求地址
// @Param method 请求方式
// @Param headers 请求header头设置
// @Param params 请求地址栏后需要拼接参数操作
// @Param data 请求报文
// @Return 返回类型 "错误信息"
func (alipayClient *DefaultAlipayClient) httpDo(url, method string, params, headers map[string]string, data []byte, alipayResponse any) (any, error) {

	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: connectTimeout,
		}).DialContext,
		ResponseHeaderTimeout: readTimeout,
	}

	//自定义client
	client := &http.Client{
		Timeout:   totalTimeout,
		Transport: trans,
	}
	//http.post等方法只是在NewRequest上又封装来了一层而已
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: "http.NewRequest is fail " + err.Error()}
	}
	req.Header.Set("Content-type", "application/json")

	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: "client.Do is fail " + err.Error()}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//通过指针将request的response值赋值，这样虽然是any类型，
	//但是我们在上次必然已经定义了any的类型
	err = json.Unmarshal(body, alipayResponse)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: "json.Unmarshal is fail " + err.Error()}
	}

	return alipayResponse, nil
}

func (alipayClient *DefaultAlipayClient) Execute(alipayRequest *request.AlipayRequest) (any, error) {
	reqPayload, err := json.Marshal(alipayRequest.Param)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: "json.Marshal is fail " + err.Error()}
	}
	path := alipayRequest.Path
	httpMethod := alipayRequest.HttpMethod
	reqTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	sign, err := genSign(fmt.Sprintf("%s", httpMethod), path, alipayClient.ClientId, reqTime, string(reqPayload), getPkcsKeu(alipayClient.MerchantPrivateKey))
	if err != nil {
		return nil, err
	}
	header := buildBaseHeader(reqTime, alipayClient.ClientId, alipayRequest.KeyVersion, sign)
	alipayResponse, err := alipayClient.httpDo(alipayClient.GatewayUrl+path, httpMethod, map[string]string{}, header, reqPayload, alipayRequest.AlipayResponse)
	if err != nil {
		return nil, err
	}
	return alipayResponse, nil
}
func getPkcsKeu(key string) string {
	return "-----BEGIN PRIVATE KEY-----\n" + key + "\n-----END PRIVATE KEY-----"
}

func genSign(httpMethod string, path string, clientId string, reqTime string, reqBody string, merchantPrivateKey string) (string, error) {
	block, _ := pem.Decode([]byte(merchantPrivateKey))
	if block == nil {
		return "", &exception.AlipayLibraryError{Message: "Failed to decode private key"}
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", &exception.AlipayLibraryError{Message: "Failed to parse private key  " + err.Error()}
	}
	payload := genSignContent(httpMethod, path, clientId, reqTime, reqBody)
	signature, err := Sign(privateKey.(*rsa.PrivateKey), []byte(payload))
	if err != nil {
		return "", &exception.AlipayLibraryError{Message: "Failed to sign data  " + err.Error()}
	}

	return signature, nil

}

func genSignContent(httpMethod string, path string, clientId string, reqTime string, reqBody string) string {
	return httpMethod + " " + path + "\n" + clientId + "." + reqTime + "." + reqBody
}

func Sign(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	// 计算数据的SHA256哈希
	hashed := sha256.Sum256(data)

	// 使用私钥签名哈希值
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.Hash.HashFunc(crypto.SHA256), hashed[:])
	if err != nil {
		return "", err
	}
	//base64编码 如果signature直接传string会造成乱码
	base64Signature := base64.StdEncoding.EncodeToString(signature)
	return Encode(base64Signature)
}

func Encode(signature string) (string, error) {
	return url.QueryEscape(signature), nil
}

func buildBaseHeader(reqTime string, clientId string, keyVersion string, signatureValue string) map[string]string {
	if keyVersion == "" {
		keyVersion = "1"
	}
	signatureValue = "algorithm=RSA256,keyVersion=" + keyVersion + ",signature=" + signatureValue
	return map[string]string{
		"Content-Type": "application/json; charset=UTF-8",
		"Request-Time": reqTime,
		"Client-Id":    clientId,
		"Key-Version":  keyVersion,
		"Signature":    signatureValue,
	}
}

func getJsonValue(jsonValue []byte, jsonField string) any {
	var jsonMap map[string]any

	if err := json.Unmarshal(jsonValue, &jsonMap); err != nil {
		log.Fatalf("json.Unmarshal is fail: %v \n")
	}

	value, ok := jsonMap[jsonField]
	if !ok {
		log.Fatalf("jsonMap is not contain jsonField: %v \n")
	}
	return value
}
