package dto

// 用户首次授权应用后，需要第二天才会产生全部的数据。​
// 只返回千粉以上用户的数据。​
// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-management/user-data/get-user-like-number

// UserDataLikeRequest 获取用户点赞数, 该接口用于获取用户点赞数。
type UserDataLikeRequest struct {
	UserDataBaseRequest
}

func (r UserDataLikeRequest) Path() string {
	return "data/external/user/like/"
}

type UserDataLike struct {
	Date    string `json:"date"`
	NewLike int64  `json:"new_like"`
}
