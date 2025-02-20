package dto

// 用户首次授权应用后，需要第二天才会产生全部的数据。​
// 只返回千粉以上用户的数据。​
// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-management/user-data/get-user-home-pv

// UserDataProfileRequest 获取用户主页访问数, 该接口用于获取用户主页访问数。​
type UserDataProfileRequest struct {
	UserDataBaseRequest
}

func (r UserDataProfileRequest) Path() string {
	return "data/external/user/profile/"
}

type UserDataProfile struct {
	Date      string `json:"date"`
	ProfileUV int64  `json:"profile_uv"`
}
