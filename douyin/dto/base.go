package dto

import (
	"fmt"
	"net/url"

	"github.com/illidaris/diplomatist/douyin/vao"
	restSender "github.com/illidaris/rest/sender"
)

type IBaseRequest interface {
	GetOptions() []restSender.Option
	Path() string
}

type BaseRequest struct {
	AccessToken string `json:"-" form:"-" url:"-"` // 抖音视频开放平台颁发的access_token
}

func (r BaseRequest) GetOptions() []restSender.Option {
	return []restSender.Option{
		restSender.WithHeader(vao.HEADER_ACCESS_TOKEN, r.AccessToken),
	}
}

func (r BaseRequest) GetUrlQuery() url.Values {
	values := url.Values{}
	values.Set("access_token", r.AccessToken)
	return values
}

type BaseExtra struct {
	BaseResult
	LogId          string `json:"log_id"`          // 标识请求的唯一id
	Now            int64  `json:"now"`             // 毫秒级时间戳
	SubDescription string `json:"sub_description"` // 子错误码描述
	SubErrorCode   int32  `json:"sub_error_code"`  // 子错误码ss
}

type BaseResult struct {
	Description string `json:"description"` // 错误码描述
	ErrorCode   int32  `json:"error_code"`  // 错误码
}

func (b BaseResult) Error() error {
	if b.ErrorCode != 0 {
		return nil
	}
	return fmt.Errorf("%s[%d]%s", vao.NAME, b.ErrorCode, b.Description)
}
