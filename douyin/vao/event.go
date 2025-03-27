package vao

type Event int32

const (
	EventNil                   Event = iota
	EventCreateVideo                 //	aweme.forward（默认开通）	用户使用开发者应用分享视频到抖音（携带分享id)
	EventAuthorize                   //	无	抖音用户授权给开发者APP
	EventUnauthorize                 //	无	抖音用户解除授权，推送事件给开发者APP
	EventImReceiveMsg                //	im.direct_message	接收私信，用户收到私信触发
	EventImSendMsg                   //	im.direct_message	发送私信，用户发送私信触发
	EventImEnterDirectMsg            //	im.direct_message	接收用户进入私信会话页事件，用户主动进入私信会话页触发
	EventImGroupReceiveMsg           //	im.group_message	接收群消息事件
	EventImGroupSendMsg              //	im.group_message	发送群消息事件
	EventEnterGroupAuditChange       //	im.group_fans.create_list 	用户加群申请
	EventGroupFansEvent              //	im.group_fans.create_list 	用户加群成功
	EventContractauthorize           //	无	用户给应用经营关系（scope名为xxx.bind，或私信相关能力）授权
	Eventcontractunauthorize         //	无	用户解除应用经营关系（scope名为xxx.bind，或私信相关能力）授权
	EventNewVideoDigg                //	user.intention	接收用户点赞事件
	EventNewFollowAction             //	user.intention	接收用户关注事件
	EventUnionAuthInfoForC           //	unionauth.identification	联合授权完成后，对c端应用进行授权信息通知
	EventUnionAuthInfoForB           //	unionauth.identification	联合授权完成后，对b端应用进行授权信息通知
)

var StringEventMap = map[string]Event{
	"create_video":             EventCreateVideo,
	"authorize":                EventAuthorize,
	"unauthorize":              EventUnauthorize,
	"im_receive_msg":           EventImReceiveMsg,
	"im_send_msg":              EventImSendMsg,
	"im_enter_direct_msg":      EventImEnterDirectMsg,
	"im_group_receive_msg":     EventImGroupReceiveMsg,
	"im_group_send_msg":        EventImGroupSendMsg,
	"enter_group_audit_change": EventEnterGroupAuditChange,
	"group_fans_event":         EventGroupFansEvent,
	"contract_authorize":       EventContractauthorize,
	"contract_unauthorize":     Eventcontractunauthorize,
	"new_video_digg":           EventNewVideoDigg,
	"new_follow_action":        EventNewFollowAction,
	"union_auth_info_for_c":    EventUnionAuthInfoForC,
	"union_auth_info_for_b":    EventUnionAuthInfoForB,
}
