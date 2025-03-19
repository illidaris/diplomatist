package agent

import (
	"context"

	"github.com/illidaris/diplomatist/douyin/dto"
)

func OAuthClientTokenInvoke(ctx context.Context, param dto.OAuthClientTokenRequest) (*dto.OAuthClientTokenResult, error) {
	api := NewPostJsonAPI[dto.OAuthClientTokenRequest, dto.OAuthClientTokenResult](param)
	return api.Invoke(ctx)
}

func OpenGetTicketInvoke(ctx context.Context, param dto.OpenGetTicketRequest) (*dto.TicketResult, error) {
	api := NewGetAPI[dto.OpenGetTicketRequest, dto.TicketResult](param)
	return api.Invoke(ctx)
}
