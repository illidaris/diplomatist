package dto

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/video-management/douyin/search-video/account-video-list

// V1VideoListRequest 搜索账号视频列表
type V1VideoListRequest struct {
	BaseRequest
	OpenId string `json:"open_id" form:"open_id" url:"open_id"` // 用户唯一标志
	Cursor int64  `json:"cursor" form:"cursor" url:"cursor"`    // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
}

func (r V1VideoListRequest) GetUrlQuery() url.Values {
	u, _ := query.Values(r)
	return u
}

func (r V1VideoListRequest) Path() string {
	return "api/douyin/v1/video/video_list/"
}

func (r V1VideoListRequest) Scope() string {
	return "video.list.bind"
}

type V1VideoListResult struct {
	BaseResult
	Cursor  int64        `json:"cursor"`
	HasMore bool         `json:"has_more"`
	List    []*VideoInfo `json:"list"`
}
