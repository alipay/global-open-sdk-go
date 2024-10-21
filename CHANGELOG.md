# Changelog


## 1.2.3 - 2024-10-21
* [#4](https://github.com/alipay/global-open-sdk-go/pull/4) feature-241022
    - add AlipayVaultingSessionRequest.is3DSAuthentication哈啰集成绑卡SDK
    - add AlipayAuthConsultRequest.authMetaData AE企业支付宝代扣集成方案
    - add Boolean  request.order.needDeclaration

## 1.2.2 - 2024-10-18
* [#4](https://github.com/alipay/global-open-sdk-go/pull/4)  update-example


## 1.2.1 - 2024-09-30
* [#3](https://github.com/alipay/global-open-sdk-go/pull/3)  update-p1
  * add MARKETPLACE - demo
  * add vaulting - demo
  * add Dispute - demo
  * add risk
  * add Notify


## 1.2.0 - 2024-09-27
* [#2](https://github.com/alipay/global-open-sdk-go/pull/2)  go-sdk-update-release
   * go-sdk-update-release

## 1.1.0 - 2024-09-27
* [#1](https://github.com/alipay/global-open-sdk-go/pull/1)  go-sdk-init
   * go-sdk-init

## 1.0.3- 2024-09-14
*  pay-update
   * update：alipayPayRequest
   * add：payments/syncArrear
   * add：payments/createDeviceCertificate

## 1.0.2 - 2024-09-14
*  add subscription example

## 1.0.1 - 2024-09-04
*  update:ScopeType
   * add TAOBAO_REBIND

## 1.0.0 - 2024-09-04
*  Init object library
    * Add pay - request response
    * Add auth - request response
    * Add customs - request response
    * Add subscription - request response
    * Add defaultAlipayClient
    * Add genSign(httpMethod string, path string, clientId string, reqTime string, reqBody string, merchantPrivateKey string) (string, error);
    * Add Sign(privateKey *rsa.PrivateKey, data []byte) (string, error);
    * Add (alipayClient *DefaultAlipayClient) Execute(alipayRequest *request.AlipayRequest) (any, error);
    * Add (alipayClient *DefaultAlipayClient) httpDo(url, method string, params, headers map[string]string, data []byte, alipayResponse any) (any, error)


