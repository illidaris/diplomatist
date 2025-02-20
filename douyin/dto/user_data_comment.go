package dto

// 用户首次授权应用后，需要第二天才会产生全部的数据。​
// 只返回千粉以上用户的数据。​
// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-management/user-data/get-user-comment-number

// UserDataCommentRequest 获取用户评论数, 该接口用于获取用户评论数。
type UserDataCommentRequest struct {
	UserDataBaseRequest
}

func (r UserDataCommentRequest) Path() string {
	return "data/external/user/comment/"
}

type UserDataComment struct {
	Date       string `json:"date"`
	NewComment int64  `json:"new_comment"`
}
