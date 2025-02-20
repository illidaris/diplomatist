package dto

// 用户首次授权应用后，需要第二天才会产生全部的数据。​
// 只返回千粉以上用户的数据。​
// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-management/user-data/get-user-video-status

// UserDataVideoRequest 获取用户视频情况, 该接口用于获取用户视频情况。
type UserDataVideoRequest struct {
	UserDataBaseRequest
}

func (r UserDataVideoRequest) Path() string {
	return "data/external/user/item/"
}

type UserDataVideo struct {
	Date       string `json:"date"`
	NewIssue   int64  `json:"new_issue"`
	NewPlay    int64  `json:"new_play"`
	TotalIssue int64  `json:"total_issue"`
}
