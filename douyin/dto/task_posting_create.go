package dto

// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/video-management/posting-task/create-posting-task
// Quota限制：10000/天
// QPS限制：100/秒
// 需要申请权限。路径：抖音开放平台控制台 > 应用详情 > 权限管理 > 能力实验室 > 投稿任务权限

// TaskPostingCreateRequest 创建投稿任务, 一个任务仅允许设置单个核销查询条件，如需查询投稿产生的多个维度数据需创建多个任务
type TaskPostingCreateRequest struct {
	BaseRequest
	StartTime     int64         `json:"start_time"`     // 开始时间秒级时间戳
	EndTime       int64         `json:"end_time"`       // 结束时间秒级时间戳
	TaskCondition TaskCondition `json:"task_condition"` // 任务条件
	TaskName      string        `json:"task_name"`      // 任务名称，长度不超过50个字符
}

func (r TaskPostingCreateRequest) Path() string {
	return "task/posting/create/"
}

func (r TaskPostingCreateRequest) Scope() string {
	return "task.posting.create"
}

// TaskPostingCreateResult 创建投稿任务结果
type TaskPostingCreateResult struct {
	TaskId     string `json:"task_id"`     // 任务id
	TaskStatus int32  `json:"task_status"` // 任务状态 1-进行中，2-未开始，3-已过期
	BaseResult
}
