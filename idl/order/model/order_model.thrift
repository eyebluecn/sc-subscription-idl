include "../enums/order_enums.thrift"
namespace go sc_subscription_api


//订单模型
struct OrderDTO {
    1: i64 id, //id
    2: i64 createTime, //创建时间
    3: i64 updateTime, //编辑时间
    4: string no, //订单唯一编号，整个系统唯一，带有前缀
    5: i64 readerId, //读者id
    6: i64 columnId, //专栏id
    7: i64 columnQuoteId, //专栏报价id
    8: i64 paymentId, //支付单id
    9: i64 price, //价格（单位：分）
    10: order_enums.OrderStatus status, //状态
}
