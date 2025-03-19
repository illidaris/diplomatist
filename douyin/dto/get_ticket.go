package dto

type TicketResult struct {
	BaseResult
	Ticket    string `json:"ticket"`     // 接口调用凭证
	ExpiresIn int64  `json:"expires_in"` // 凭证有效时间，单位：秒。目前是7200秒之内的值。
}
