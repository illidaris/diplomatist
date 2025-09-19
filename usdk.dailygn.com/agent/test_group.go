package agent

import (
	"encoding/json"

	dto "github.com/illidaris/diplomatist/usdk.dailygn.com/dto"
	restSender "github.com/illidaris/rest/sender"
)

var _ = restSender.IRequest(&TestGroupListAPI{})

func NewTestGroupListAPI(param dto.TestGroupInfoListRequest) *TestGroupListAPI {
	return &TestGroupListAPI{
		TestGroupInfoListRequest: param,
		Response:                 &dto.TestGroupInfoListResponse{},
	}
}

type TestGroupListAPI struct {
	restSender.GETRequest `json:"-"`
	dto.TestGroupInfoListRequest
	Response *dto.TestGroupInfoListResponse `json:"-"`
}

func (r *TestGroupListAPI) GetResponse() interface{} {
	return r.Response
}

func (r *TestGroupListAPI) Decode(bs []byte) error {
	if r.Response == nil {
		r.Response = &dto.TestGroupInfoListResponse{}
	}
	return json.Unmarshal(bs, r.Response)
}
