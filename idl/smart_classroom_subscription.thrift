include "./subscription/subscription_api.thrift"
include "./mq/mq_api.thrift"
namespace go sc_subscription_api

//定义subscription应用中的所有RPC服务
service SubscriptionService {

    //查询订阅列表
    subscription_api.SubscriptionPageResponse SubscriptionPage(1: subscription_api.SubscriptionPageRequest request)

    //准备订阅
    subscription_api.SubscriptionPrepareResponse SubscriptionPrepare(1: subscription_api.SubscriptionPrepareRequest request)

    //mq消息到达
    mq_api.MqMessageArriveResponse MqMessageArrive(1: mq_api.MqMessageArriveRequest request)

}
