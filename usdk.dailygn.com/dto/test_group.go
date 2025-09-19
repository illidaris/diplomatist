package dto

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

type TestGroupInfoListRequest struct {
	BaseDateRequest
}

func (r TestGroupInfoListRequest) GetUrlQuery() url.Values {
	v, _ := query.Values(r)
	return v
}

func (r *TestGroupInfoListRequest) Encode() ([]byte, error) {
	r.Body = ""
	return nil, nil
}

func (r TestGroupInfoListRequest) GetAction() string {
	return "webcast/gamecp/attribute/channel/get_test_group_infos"
}

type TestGroupInfoListResponse struct {
	BaseResponse
	Data []TestGroupInfo `json:"test_group_info_list"`
}

type TestGroupInfo struct {
	GroupId              string `json:"group_id"`                 // 实验组ID
	GroupName            string `json:"group_name"`               // 实验组名称
	TestGameUserIdAmount int64  `json:"test_game_user_id_amount"` // 实验组官包用户总数
	TestOpenIdAmount     int64  `json:"test_open_id_amount"`      // 实验组流水分账包用户总数
}
