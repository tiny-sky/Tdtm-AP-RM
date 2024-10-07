package ap

import (
	"github.com/tiny-sky/Tdtm-Client/RM/account"
	"github.com/tiny-sky/Tdtm-Client/RM/order"
	"github.com/tiny-sky/Tdtm-Client/RM/stock"
	"github.com/tiny-sky/Tdtm/client"
	"github.com/tiny-sky/Tdtm/core/registry"
	"github.com/tiny-sky/Tdtm/core/resolver"
)

func GetSrv() []*client.Group {
	m := client.NewManger()

	m.AddGroups(client.NewTccGroup(
		"http://127.0.0.1:10001/order/try",
		"http://127.0.0.1:10001/order/confirm",
		"http://127.0.0.1:10001/order/cancel").SetData(order.NewData()).
		SetTimeout(2))

	m.AddGroups(client.NewTccGroup(
		"http://127.0.0.1:10002/stock/try",
		"http://127.0.0.1:10002/stock/confirm",
		"http://127.0.0.1:10002/stock/cancel").SetData(stock.NewData()).
		SetTimeout(2))

	m.AddGroups(client.NewTccGroup(
		"http://127.0.0.1:10003/account/try",
		"http://127.0.0.1:10003/account/confirm",
		"http://127.0.0.1:10003/account/cancel").SetData(account.NewData()).
		SetTimeout(2))
	return m.Groups()
}

func RegisterBuilder(discovery registry.Discovery) {
	resolver.Register(discovery)
}
