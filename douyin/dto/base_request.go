package dto

import (
	"net/url"

	"github.com/illidaris/diplomatist/douyin/vao"
	restSender "github.com/illidaris/rest/sender"
)

type IBaseRequest interface {
	GetOptions() []restSender.Option
	Path() string
}
type IErorResponse interface {
	Error() error
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
	values.Set(vao.HEADER_ACCESS_TOKEN, r.AccessToken)
	return values
}
