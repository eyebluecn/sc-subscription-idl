include "../base/base.thrift"
include "../base/pagination.thrift"
include "enums/subscription_enums.thrift"
include "model/subscription_model.thrift"
include "../order/model/order_model.thrift"
namespace go sc_subscription_api

//订阅列表 请求体
struct SubscriptionPageRequest {
	1: i64 pageNum //当前页码 1基
	2: i64 pageSize //每页大小
	3: optional i64 readerId //读者id
	4: optional list<i64> columnIds //专栏id
	5: optional i64 orderId //订单id
	6: optional subscription_enums.SubscriptionStatus status //状态

	255: optional base.Base base //标准请求内容
}

//订阅列表 响应体
struct SubscriptionPageResponse {
	1: list<subscription_model.SubscriptionDTO> data //数据信息
    2: pagination.Pagination pagination //分页指示器

    255: base.BaseResp baseResp //标准返回内容

}

//准备订阅 请求体
struct SubscriptionPrepareRequest {
	1: i64 columnId //专栏
	2: string payMethod //支付方式
	3: i64 readerId //当前操作的读者id

	255: optional base.Base base //标准请求内容
}

//准备订阅 响应体
struct SubscriptionPrepareResponse {
	1: SubscriptionPrepareData data //数据信息

    255: base.BaseResp baseResp //标准返回内容
}

struct SubscriptionPrepareData {
	1: order_model.OrderDTO orderDTO //订单
	2: string thirdTransactionNo //支付的一些token及信息
	3: string nonceStr //支付时候的验证信息等
}
