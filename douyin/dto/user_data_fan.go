package dto

// 用户首次授权应用后，需要第二天才会产生全部的数据。​
// 只返回千粉以上用户的数据。​
// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-management/user-data/get-user-fans-count

// UserDataFansRequest 获取用户粉丝数情况, 该接口用于获取用户粉丝数。
type UserDataFansRequest struct {
	UserDataBaseRequest
}

func (r UserDataFansRequest) Path() string {
	return "data/external/user/fans/"
}

type UserDataFans struct {
	Date      string `json:"date"`
	NewFans   int64  `json:"new_fans"`
	TotalFans int64  `json:"total_fans"`
}
