package dto

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"time"

	restSender "github.com/illidaris/rest/sender"
	"github.com/illidaris/rest/signature"
	"github.com/spf13/cast"
)

type IRequest interface {
	restSender.IRequest
	GetSigns(method, path string, query url.Values) string
}

type BaseRequest struct {
	AppId     string `json:"app_id" form:"app_id" url:"app_id"`
	AppSecret string `json:"-" form:"-" url:"-"`
	Ver       string `json:"-" form:"-" url:"-"`
	Alg       string `json:"-" form:"-" url:"-"`
	Nonce     string `json:"-" form:"-" url:"-"`
	Ts        int64  `json:"-" form:"-" url:"-"`
	Sign      string `json:"-" form:"-" url:"-"`
	Body      string `json:"-" form:"-" url:"-"`
}

func (r *BaseRequest) GetSigns(method, path string, query url.Values) string {
	if r.Ts == 0 {
		r.Ts = time.Now().UnixMilli()
	}
	if r.Nonce == "" {
		r.Nonce = signature.DefaultNoiseRand()
	}
	if r.Ver == "" {
		r.Ver = "1.0"
	}
	if r.Alg == "" {
		r.Alg = "HMAC-SHA256"
	}
	path = "/" + path
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	words := []string{
		method,
		path,
		cast.ToString(r.Ts),
		r.Nonce,
	}
	words = append(words, r.Body)
	r.Sign = HashMac(sha256.New, r.AppSecret, words...)
	res := fmt.Sprintf("appid=\"%s\",v=\"%s\",alg=\"%s\",nonce=\"%s\",timestamp=\"%s\",signature=\"%s\"",
		r.AppId,
		r.Ver,
		r.Alg,
		r.Nonce,
		cast.ToString(r.Ts),
		r.Sign)
	return res
}

type BaseDateRequest struct {
	BaseRequest
	Date string `json:"date" form:"date" url:"date"`
}
