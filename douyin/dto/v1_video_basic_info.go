package dto

// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/video-management/posting-task/video-basic-info
// 当 账号为“私密账号” 或 视频为“私密视频/部分好友可见/密友可见/朋友可见” 或 视频已删除，视频会被认为非公开
// 需要申请权限。路径：抖音开放平台控制台 > 应用详情 > 权限管理 > 能力实验室 > 用户投稿任务

// V1VideoBasicInfoRequest 查询视频基础信息
type V1VideoBasicInfoRequest struct {
	BaseRequest
	ItemIds  []string `json:"item_ids"`  // item_id数组，仅能查询access_token对应用户上传的视频
	VideoIds []string `json:"video_ids"` // item_id数组，明文item_id，如果和item_ids同时传入，优先处理video_ids
}

func (r V1VideoBasicInfoRequest) Path() string {
	return "api/douyin/v1/video/video_basic_info/"
}

func (r V1VideoBasicInfoRequest) Scope() string {
	return "posting.behavior"
}

type V1VideoBasicInfoResult struct {
	BaseResult
	List []*VideoBasicInfo `json:"list"`
}
