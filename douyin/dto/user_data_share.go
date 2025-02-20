package dto

// 用户首次授权应用后，需要第二天才会产生全部的数据。​
// 只返回千粉以上用户的数据。​
// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-management/user-data/get-user-share-count

// UserDataShareRequest 获取用户分享数, 该接口用于获取用户分享数。
type UserDataShareRequest struct {
	UserDataBaseRequest
}

func (r UserDataShareRequest) Path() string {
	return "data/external/user/share/"
}

type UserDataShare struct {
	Date     string `json:"date"`
	NewShare int64  `json:"new_share"`
}
