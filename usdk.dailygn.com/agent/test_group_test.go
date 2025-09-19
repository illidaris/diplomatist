package agent

import (
	"context"
	"testing"

	"github.com/illidaris/core"
	"github.com/illidaris/diplomatist/usdk.dailygn.com/dto"
)

func TestNewTestGroupListAPI(t *testing.T) {
	ctx := context.Background()
	param := dto.TestGroupInfoListRequest{}
	param.Date = "20250919"
	param.AppId = ""
	param.Ts = 1758264672826
	param.Nonce = "rVhSUz"
	param.AppSecret = ""

	api := NewTestGroupListAPI(param)
	ctx = context.WithValue(ctx, core.TraceID, "tttttrace12333")

	err := SignInvoke(ctx, api, "")
	if err != nil {
		t.Error(err)
	}
}

func TestNewTestGroupDataListAPI(t *testing.T) {
	ctx := context.Background()
	param := dto.TestGroupDataListRequest{}
	param.Date = "20250919"
	param.AppId = ""
	param.Ts = 1758264672826
	param.Nonce = "rVhSUz"
	param.AppSecret = ""
	param.PageSize = 1
	param.PageOffset = 0
	param.UserType = 0
	param.GroupId = "test_group_id"

	api := NewTestGroupDataListAPI(param)
	ctx = context.WithValue(ctx, core.TraceID, "tttttrace12333")

	err := SignInvoke(ctx, api, "")
	if err != nil {
		t.Error(err)
	}
}
