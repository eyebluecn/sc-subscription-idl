include "../../subscription/model/subscription_model.thrift"
include "../../column/model/column_model.thrift"
include "../../reader/model/reader_model.thrift"
include "../../order/model/order_model.thrift"

namespace go sc_subscription_api



//富订阅模型
struct RichSubscriptionDTO {
    1: subscription_model.SubscriptionDTO subscription //订阅本身
    2: column_model.ColumnDTO column //专栏
    3: reader_model.ReaderDTO reader //读者
    4: order_model.OrderDTO order //订单
}