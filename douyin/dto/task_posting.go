package dto

// TaskCondition 任务条件
type TaskCondition struct {
	Condition string `json:"condition"` // 收藏:collection, 评论:comment, 点赞:digg, 播放:play, 分享:share, 下载:download
	MaxValue  int64  `json:"max_value"` // 阈值 右区间, 播放数:最大为1000w, 其他:最大10w
	MinValue  int64  `json:"min_value"` // 阈值 左区间, 最小为0
}

type TaskBaseResponse[T any] struct {
	Data  *T         `json:"data"`
	Extra *BaseExtra `json:"extra"`
}
