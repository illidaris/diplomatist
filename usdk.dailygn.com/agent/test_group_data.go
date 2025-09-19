package agent

import (
	"encoding/json"

	dto "github.com/illidaris/diplomatist/usdk.dailygn.com/dto"
	restSender "github.com/illidaris/rest/sender"
)

var _ = restSender.IRequest(&TestGroupListAPI{})

func NewTestGroupDataListAPI(param dto.TestGroupDataListRequest) *TestGroupDataListAPI {
	return &TestGroupDataListAPI{
		TestGroupDataListRequest: param,
		Response:                 &dto.TestGroupDataListResponse{},
	}
}

type TestGroupDataListAPI struct {
	restSender.GETRequest `json:"-"`
	dto.TestGroupDataListRequest
	Response *dto.TestGroupDataListResponse `json:"-"`
}

func (r *TestGroupDataListAPI) GetResponse() interface{} {
	return r.Response
}

func (r *TestGroupDataListAPI) Decode(bs []byte) error {
	if r.Response == nil {
		r.Response = &dto.TestGroupDataListResponse{}
	}
	return json.Unmarshal(bs, r.Response)
}
