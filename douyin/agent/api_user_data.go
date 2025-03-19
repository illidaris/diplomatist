package agent

import (
	"context"

	"github.com/illidaris/diplomatist/douyin/dto"
)

func UserDataVideoInvoke(ctx context.Context, param dto.UserDataVideoRequest) (*dto.UserDataBaseResult[dto.UserDataVideo], error) {
	api := NewUserDataAPI[dto.UserDataVideoRequest, dto.UserDataVideo](param)
	return api.Invoke(ctx)
}

func UserDataLikeInvoke(ctx context.Context, param dto.UserDataLikeRequest) (*dto.UserDataBaseResult[dto.UserDataLike], error) {
	api := NewUserDataAPI[dto.UserDataLikeRequest, dto.UserDataLike](param)
	return api.Invoke(ctx)
}

func UserDataFansInvoke(ctx context.Context, param dto.UserDataFansRequest) (*dto.UserDataBaseResult[dto.UserDataFans], error) {
	api := NewUserDataAPI[dto.UserDataFansRequest, dto.UserDataFans](param)
	return api.Invoke(ctx)
}

func UserDataCommentInvoke(ctx context.Context, param dto.UserDataCommentRequest) (*dto.UserDataBaseResult[dto.UserDataComment], error) {
	api := NewUserDataAPI[dto.UserDataCommentRequest, dto.UserDataComment](param)
	return api.Invoke(ctx)
}

func UserDataShareInvoke(ctx context.Context, param dto.UserDataShareRequest) (*dto.UserDataBaseResult[dto.UserDataShare], error) {
	api := NewUserDataAPI[dto.UserDataShareRequest, dto.UserDataShare](param)
	return api.Invoke(ctx)
}

func UserDataProfileInvoke(ctx context.Context, param dto.UserDataProfileRequest) (*dto.UserDataBaseResult[dto.UserDataProfile], error) {
	api := NewUserDataAPI[dto.UserDataProfileRequest, dto.UserDataProfile](param)
	return api.Invoke(ctx)
}
