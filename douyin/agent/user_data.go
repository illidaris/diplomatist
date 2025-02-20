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

func NewUserDataAPI[Req dto.IBaseRequest, T any](param Req) *UserDataBaseAPI[Req, T] {
	return &UserDataBaseAPI[Req, T]{
		Request:  param,
		Response: &dto.UserDataBaseResponse[T]{},
	}
}

var _ = restSender.IRequest(&UserDataBaseAPI[dto.IBaseRequest, any]{})

type UserDataBaseAPI[Req dto.IBaseRequest, T any] struct {
	restSender.GETRequest `json:"-"`
	Request               Req
	Response              *dto.UserDataBaseResponse[T] `json:"-"`
}

func (r UserDataBaseAPI[Req, T]) GetUrlQuery() url.Values {
	u, _ := query.Values(r.Request)
	return u
}

func (r UserDataBaseAPI[Req, T]) Encode() ([]byte, error) {
	return nil, nil
}

func (r UserDataBaseAPI[Req, T]) GetResponse() interface{} {
	return r.Response
}

func (r UserDataBaseAPI[Req, T]) GetAction() string {
	return r.Request.Path()
}

func (r UserDataBaseAPI[Req, T]) Decode(bs []byte) error {
	if r.Response == nil {
		r.Response = &dto.UserDataBaseResponse[T]{}
	}
	return json.Unmarshal(bs, r.Response)
}

func (r *UserDataBaseAPI[Req, T]) Invoke(ctx context.Context, opts ...restSender.Option) (*dto.UserDataBaseResult[T], error) {
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
	if r.Response.Data == nil {
		return nil, vao.ErrResponseNil
	}
	if r.Response.Data.Error() != nil {
		return r.Response.Data, r.Response.Data.Error()
	}
	return r.Response.Data, nil
}
