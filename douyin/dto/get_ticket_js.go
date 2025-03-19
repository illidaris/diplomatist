package dto

// https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/tools-ability/jsb-management/get-jsticket
// 该接口用于获取 jsapi_ticket；本接口适用于抖音
// jsapi_ticket（同一个ticket的有效期大约在7200秒左右，开发者需要根据接口实际返回的过期时间为准）

type JsGetTicketRequest struct {
	BaseRequest // client_token
}

func (r JsGetTicketRequest) Path() string {
	return "js/getticket/"
}
