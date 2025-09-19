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
	param.Date = "20240730"
	param.AppId = ""
	param.AppSecret = ""

	api := NewTestGroupListAPI(param)
	ctx = context.WithValue(ctx, core.TraceID, "tttttrace12333")

	err := SignInvoke(ctx, api, "")
	if err != nil {
		t.Error(err)
	}
}
