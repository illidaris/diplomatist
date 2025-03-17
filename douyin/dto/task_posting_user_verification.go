package dto

// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/video-management/posting-task/verify-posting-task
// 该接口用于核销投稿任务
// Quota限制：无
// QPS限制：100/秒
// 移动/网站应用：抖音开放平台控制台 > 应用详情 > 权限管理 > 能力实验室 > 投稿任务权限

// TaskPostingUserVerificationRequest 核销投稿任务
type TaskPostingUserVerificationRequest struct {
	BaseRequest
	TargetOpenId string `json:"target_open_id"` // 用户唯一标志
	TaskID       string `json:"task_id"`        // 任务ID
	VideoID      string `json:"video_id"`       // 视频ID
}

func (r TaskPostingUserVerificationRequest) Path() string {
	return "task/posting/user/"
}

func (r TaskPostingUserVerificationRequest) Scope() string {
	return "task.posting.user_verification"
}

// TaskPostingUserVerificationResult 核销投稿任务结果
type TaskPostingUserVerificationResult struct {
	Result bool `json:"result"` // 结果
	BaseResult
}
