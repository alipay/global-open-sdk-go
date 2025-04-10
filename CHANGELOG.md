# Changelog

## 1.2.12 - 2025-04-10
* [#14](https://github.com/alipay/global-open-sdk-go/pull/14) feature-250410
    - Antom 印度渠道接入AMS拒付相关接口的标准变更
    - Antom印度渠道接入（UPI/CARD/NETBAKING）相关接口标准变更
    - CKP二期支持商户传入支付方式地区和支付方式要素

## 1.2.11 - 2025-02-05
* [#13](https://github.com/alipay/global-open-sdk-go/pull/13) feature-250205
    - 支付、查询、支付结果通知新增卡相关信息字段

## 1.2.10 - 2025-01-22
* [#12](https://github.com/alipay/global-open-sdk-go/pull/12) feature-250122
    - update promotionResults


## 1.2.9 - 2025-01-06
* [#11](https://github.com/alipay/global-open-sdk-go/pull/11) feature-250106
    - 订阅支付新增“更新接口”
    - 增加验签

## 1.2.8 - 2024-12-24
* [#10](https://github.com/alipay/global-open-sdk-go/pull/10) feature-241224
    - CKP二期支持商户传入可选支付方式列表
    - AMS独立绑卡支持MIT交易


## 1.2.7 - 2024-12-16
* [#9](https://github.com/alipay/global-open-sdk-go/pull/9) feature-241216
    - RDR拒付通知优化通用能力变更
    - Antom新增ApplePay支付方式

## 1.2.6 - 2024-12-02
* [#8](https://github.com/alipay/global-open-sdk-go/pull/8) feature-241202
    - update AlipayDisputeNotify
    - update AlipayPayQueryResponse

## 1.2.5 - 2024-11-25
* [#7](https://github.com/alipay/global-open-sdk-go/pull/7) feature-241125
    - update AlipayPayResultNotify
    - update AlipayPayQueryRequest

## 1.2.5 - 2024-11-25
* [#6](https://github.com/alipay/global-open-sdk-go/pull/6) feature-241125
    - update Leg

## 1.2.4 - 2024-10-24
* [#5](https://github.com/alipay/global-open-sdk-go/pull/5) feature-notify
    - update notify

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


