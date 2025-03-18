package dto

type VideoBasicInfo struct {
	Cover       string `json:"cover"`        // 视频封面
	CreateTime  int64  `json:"create_time"`  // 视频创建时间戳
	ItemId      string `json:"item_id"`      // 视频id
	MediaType   int    `json:"media_type"`   // 媒体类型。2:图集;4:视频
	Title       string `json:"title"`        // 视频标题
	VideoId     string `json:"video_id"`     // 视频真实id
	VideoStatus int    `json:"video_status"` // 视频状态
}

type VideoInfo struct {
	VideoBasicInfo
	Statistics VideoStatistics `json:"statistics"`  // 视频统计信息
	IsReviewed bool            `json:"is_reviewed"` // 表示是否审核结束。审核通过或者失败都会返回true，审核中返回false。
	IsTop      bool            `json:"is_top"`      // 是否置顶
	ShareUrl   string          `json:"share_url"`   // 视频播放页面。视频播放页可能会失效，请在观看视频前调用/video/data/获取最新的播放页。
}

type VideoStatistics struct {
	CommentCount  int32 `json:"comment_count"`  // 评论数
	DiggCount     int32 `json:"digg_count"`     // 点赞数
	DownloadCount int32 `json:"download_count"` // 下载数
	ForwardCount  int32 `json:"forward_count"`  // 转发数
	PlayCount     int32 `json:"play_count"`     // 播放数 播放数，只有作者本人可见。公开视频设为私密后，播放数也会返回0。
	ShareCount    int32 `json:"share_count"`    // 分享数
}
