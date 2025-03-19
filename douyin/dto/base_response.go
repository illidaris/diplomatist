package dto

import (
	"fmt"

	"github.com/illidaris/diplomatist/douyin/vao"
)

type BaseResponse[T any] struct {
	Data  *T         `json:"data"`
	Extra *BaseExtra `json:"extra"`
}

func (r BaseResponse[T]) Error() error {
	if r.Extra != nil && r.Extra.Error() != nil {
		return r.Extra.Error()
	}
	if r.Data == nil {
		return vao.ErrResponseNil
	}
	if errData, ok := any(r.Data).(IErorResponse); ok {
		if errData.Error() != nil {
			return errData.Error()
		}
	}
	return nil
}

type BaseExtra struct {
	BaseResult
	LogId          string `json:"log_id"`          // 标识请求的唯一id
	Now            int64  `json:"now"`             // 毫秒级时间戳
	SubDescription string `json:"sub_description"` // 子错误码描述
	SubErrorCode   int32  `json:"sub_error_code"`  // 子错误码ss
}

type BaseResult struct {
	Description string `json:"description"` // 错误码描述
	ErrorCode   int32  `json:"error_code"`  // 错误码
}

func (b BaseResult) Error() error {
	if b.ErrorCode == 0 {
		return nil
	}
	return fmt.Errorf("%s[%d]%s", vao.NAME, b.ErrorCode, b.Description)
}
