package dto

import restSender "github.com/illidaris/rest/sender"

// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-permission/client-token
// 该接口用于获取接口调用的凭证 client_token。该接口适用于抖音授权；
// client_token 用于不需要用户授权就可以调用的接口。
// client_token 的有效时间为 2 个小时，重复获取 client_token 后会使上次的 client_token 失效（但有 5 分钟的缓冲时间，连续多次获取 client_token 只会保留最新的两个 client_token）。

func NewOAuthClientTokenRequest(key, secret string) OAuthClientTokenRequest {
	return OAuthClientTokenRequest{
		GrantType:    "client_credential",
		ClientKey:    key,
		ClientSecret: secret,
	}
}

type OAuthClientTokenRequest struct {
	GrantType    string `json:"grant_type"`    // 必须，请求参数，填写 client_credential
	ClientKey    string `json:"client_key"`    // 必须，应用唯一标识，在开放平台申请创建应用后生成
	ClientSecret string `json:"client_secret"` // 必须，应用密钥，在开放平台申请创建应用后生成
}

func (r OAuthClientTokenRequest) Path() string {
	return "oauth/client_token/"
}
func (r OAuthClientTokenRequest) GetOptions() []restSender.Option {
	return []restSender.Option{}
}

type OAuthClientTokenResult struct {
	BaseResult
	AccessToken string `json:"access_token"` // 接口调用凭证
	ExpiresIn   int64  `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
}
