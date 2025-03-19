package dto

// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/tools-ability/jsb-management/open-ticket
// 需要申请权限，不需要用户授权
// 该接口用于 h5 链接拉起抖音发布器分享视频时对开发者身份进行验签；本接口适用于抖音。
// 抖音的 OAuth API 以https://open.douyin.com/开头。
// open_ticket 的有效时间为 2 个小时，重复获取 ticket 后会使上次的 ticket 失效（但有 5 分钟的缓冲时间）。

type OpenGetTicketRequest struct {
	BaseRequest // client_token
}

func (r OpenGetTicketRequest) Path() string {
	return "open/getticket/"
}
