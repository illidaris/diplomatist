package agent

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/illidaris/diplomatist/douyin/dto"
	"github.com/illidaris/diplomatist/douyin/vao"
	restSender "github.com/illidaris/rest/sender"
)

func NewPostJsonAPI[Req dto.IBaseRequest, T any](param Req) *PostJsonBaseAPI[Req, T] {
	return &PostJsonBaseAPI[Req, T]{
		Request:  param,
		Response: &dto.BaseResponse[T]{},
	}
}

var _ = restSender.IRequest(&PostJsonBaseAPI[dto.IBaseRequest, any]{})

type PostJsonBaseAPI[Req dto.IBaseRequest, T any] struct {
	restSender.JSONRequest `json:"-"`
	Request                Req
	Response               *dto.BaseResponse[T] `json:"-"`
}

func (r PostJsonBaseAPI[Req, T]) GetUrlQuery() url.Values {
	u, _ := query.Values(r.Request)
	return u
}

func (r PostJsonBaseAPI[Req, T]) Encode() ([]byte, error) {
	return json.Marshal(r.Request)
}

func (r PostJsonBaseAPI[Req, T]) GetResponse() interface{} {
	return r.Response
}

func (r PostJsonBaseAPI[Req, T]) GetAction() string {
	return r.Request.Path()
}

func (r PostJsonBaseAPI[Req, T]) Decode(bs []byte) error {
	if r.Response == nil {
		r.Response = &dto.BaseResponse[T]{}
	}
	return json.Unmarshal(bs, r.Response)
}

func (r *PostJsonBaseAPI[Req, T]) Invoke(ctx context.Context, opts ...restSender.Option) (*T, error) {
	if len(opts) == 0 {
		opts = []restSender.Option{
			restSender.WithHost(vao.BASE_HOST),
			restSender.WithTimeConsume(true),
			restSender.WithTimeout(time.Second * 10),
		}
	}
	opts = append(opts, r.Request.GetOptions()...)
	_, err := restSender.NewSender(opts...).Invoke(ctx, r)
	if err != nil {
		return nil, err
	}
	return r.Response.Data, r.Response.Error()
}
