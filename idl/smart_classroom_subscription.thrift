include "./subscription/subscription_api.thrift"
namespace go sc_subscription_api

//定义subscription应用中的所有RPC服务
service SubscriptionService {

    //查询订阅列表
    subscription_api.SubscriptionListResponse SubscriptionList(1: subscription_api.SubscriptionListRequest request)

}
