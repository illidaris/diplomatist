package dto

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

type TestGroupDataListRequest struct {
	BaseDateRequest
	UserType   int64  `json:"user_type" form:"user_type" url:"user_type"`       // 0:官包对应的game_user_id 1:抖音流水分账包对应的sdk_open_id
	GroupId    string `json:"group_id" form:"group_id" url:"group_id"`          // 需要查询实验信息的日期，该时间为需要统计汇总付费的日期，刷新时间一般为T+1，日期格式：20240730
	PageSize   int64  `json:"page_size" form:"page_size" url:"page_size"`       // 分页大小，最大值为10000，如果失败请求过多建议调小
	PageOffset int64  `json:"page_offset" form:"page_offset" url:"page_offset"` // 分页下标，起始值为0
}

func (r TestGroupDataListRequest) GetUrlQuery() url.Values {
	v, _ := query.Values(r)
	return v
}

func (r *TestGroupDataListRequest) Encode() ([]byte, error) {
	r.Body = ""
	return nil, nil
}

func (r TestGroupDataListRequest) GetAction() string {
	return "webcast/gamecp/attribute/channel/get_test_group_data"
}

type TestGroupDataListResponse struct {
	BaseResponse
	Data []string `json:"user_id_list"`
}
