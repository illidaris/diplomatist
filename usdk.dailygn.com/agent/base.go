package agent

import (
	"context"
	"time"

	dto "github.com/illidaris/diplomatist/usdk.dailygn.com/dto"
	restSender "github.com/illidaris/rest/sender"
)

var (
	defaultHost = "https://usdk.dailygn.com"
)

// Invoke API执行
func SignInvoke(ctx context.Context, req dto.IRequest, host string) error {
	if host == "" {
		host = defaultHost
	}
	v := req.GetSigns(req.GetMethod(), req.GetAction(), req.GetUrlQuery())
	s := restSender.NewSender(
		restSender.WithTimeout(time.Second*20),
		restSender.WithHeader("x-game-sign", v),
		restSender.WithHost(host),
	)
	_, err := s.Invoke(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
