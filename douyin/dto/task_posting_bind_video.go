package dto

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/video-management/posting-task/bind-video
// 视频作者必须为OpenID对应的用户
// 视频必须为公开状态当 账号为“私密账号” 或 视频为“私密视频/部分好友可见/密友可见/朋友可见” 或 视频已删除，视频会被认为非公开
// 视频必须为通过ClientToken对应的应用发布
// 视频发布时间不得早于任务开始时间
// 视频必须为视频/图集类型
// Quota限制：无
// QPS限制：1000/秒

// TaskPostingBindVideoRequest 投稿任务视频绑定
type TaskPostingBindVideoRequest struct {
	BaseRequest
	OpenId  string `json:"-" form:"open_id" url:"open_id"` // 用户唯一标志
	TaskID  string `json:"task_id"`                        // 创建任务之后获取的任务ID
	VideoID string `json:"video_id"`                       // 视频ID
}

func (r TaskPostingBindVideoRequest) GetUrlQuery() url.Values {
	u, _ := query.Values(r)
	return u
}

func (r TaskPostingBindVideoRequest) Path() string {
	return "task/posting/bind_video/"
}

func (r TaskPostingBindVideoRequest) Scope() string {
	return "posting.behavior"
}

// TaskPostingBindVideoResult 投稿任务绑定视频结果
type TaskPostingBindVideoResult struct {
	BaseResult
}
