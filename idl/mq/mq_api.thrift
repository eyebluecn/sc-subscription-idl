include "../base/base.thrift"
include "../base/pagination.thrift"

namespace go sc_subscription_api

// 请求体
struct MqMessageArriveRequest {
	1: string topic //消息TOPC
	2: string tags //消息的tags
	3: string keys //消息的Key字段是为了唯一标识消息的，方便运维排查问题。如果不设置Key，则无法定位消息丢失原因。
	4: string body //发送的消息体，在智慧课堂中只发送文本消息。

	255: optional base.Base base //标准请求内容
}

//订阅列表 响应体
struct MqMessageArriveResponse {

    255: base.BaseResp baseResp //标准返回内容
}

