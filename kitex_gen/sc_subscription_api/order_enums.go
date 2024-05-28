// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package sc_subscription_api

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type OrderStatus int64

const (
	OrderStatus_CREATED    OrderStatus = 0
	OrderStatus_PAID       OrderStatus = 1
	OrderStatus_SUBSCRIBED OrderStatus = 2
	OrderStatus_CLOSED     OrderStatus = 3
	OrderStatus_CANCELED   OrderStatus = 4
)

func (p OrderStatus) String() string {
	switch p {
	case OrderStatus_CREATED:
		return "CREATED"
	case OrderStatus_PAID:
		return "PAID"
	case OrderStatus_SUBSCRIBED:
		return "SUBSCRIBED"
	case OrderStatus_CLOSED:
		return "CLOSED"
	case OrderStatus_CANCELED:
		return "CANCELED"
	}
	return "<UNSET>"
}

func OrderStatusFromString(s string) (OrderStatus, error) {
	switch s {
	case "CREATED":
		return OrderStatus_CREATED, nil
	case "PAID":
		return OrderStatus_PAID, nil
	case "SUBSCRIBED":
		return OrderStatus_SUBSCRIBED, nil
	case "CLOSED":
		return OrderStatus_CLOSED, nil
	case "CANCELED":
		return OrderStatus_CANCELED, nil
	}
	return OrderStatus(0), fmt.Errorf("not a valid OrderStatus string")
}

func OrderStatusPtr(v OrderStatus) *OrderStatus { return &v }
func (p *OrderStatus) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = OrderStatus(result.Int64)
	return
}

func (p *OrderStatus) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}
