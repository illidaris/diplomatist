package dto

type UserDataBaseRequest struct {
	BaseRequest
	OpenId   string `json:"open_id,omitempty" form:"open_id,omitempty" url:"open_id,omitempty"`       // 用户唯一标志，通过 /oauth/access_token 返回的 open_id 字段获取
	DataType int64  `json:"data_type,omitempty" form:"data_type,omitempty" url:"data_type,omitempty"` // 近7/15/30天；输入7代表7天、15代表15天、30代表30天。
}

func (r UserDataVideoRequest) Scope() string {
	return "data.external.user"
}

type UserDataBaseResponse[T any] struct {
	Data  *UserDataBaseResult[T] `json:"data"`
	Extra *BaseExtra             `json:"extra"`
}

type UserDataBaseResult[T any] struct {
	BaseResult
	ResultList []*T `json:"result_list"`
}
