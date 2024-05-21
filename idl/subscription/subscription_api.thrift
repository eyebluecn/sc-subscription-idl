include "../base/base.thrift"
include "../base/pagination.thrift"
include "enums/enums.thrift"
include "model/model.thrift"
namespace go sc_subscription_api

//订阅列表 请求体
struct SubscriptionPageRequest {
	1: i64 pageNum //当前页码 1基
	2: i64 pageSize //每页大小
	3: optional i64 readerId //读者id
	4: optional list<i64> columnIds //专栏id
	5: optional i64 orderId //订单id
	6: optional enums.SubscriptionStatus status //状态

	255: optional base.Base base //标准请求内容
}

//订阅列表 响应体
struct SubscriptionPageResponse {
	1: list<model.SubscriptionDTO> data //数据信息
    2: pagination.Pagination pagination //分页指示器

    255: base.BaseResp baseResp //标准返回内容

}
