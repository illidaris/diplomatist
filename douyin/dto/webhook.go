package dto

import (
	"crypto/sha1"
	"fmt"
)

type WebhookRequest struct {
	Event      string `json:"event"`
	ClientKey  string `json:"client_key"`
	FromUserId string `json:"from_user_id"`
	ToUserId   string `json:"to_user_id"`
	LogId      string `json:"logid"`
	Content    any    `json:"content"`
}

// DebugVerifyMsg 调试验证消息
type DebugVerifyMsg struct {
	Challenge int64 `json:"challenge"`
}

type CreateVideoMsg struct {
	ItemId  string `json:"item_id"`  // 创建的视频id
	ShareId string `json:"share_id"` // 标识分享的share_id
}

type AuthorizeMsg struct {
	Scopes []string `json:"scopes"` // 授权scope列表
}

// DebugVerifyResponse 调试验证响应
type DebugVerifyResponse struct {
	DebugVerifyMsg
}

func SignSignature(body []byte, secret string) string {
	h := sha1.New()
	h.Write([]byte(secret))
	h.Write(body)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
