package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/tiny-sky/Tdtm-Client/RM/account"
	"github.com/tiny-sky/Tdtm-Client/RM/conf"
	"github.com/tiny-sky/Tdtm-Client/RM/order"
	"github.com/tiny-sky/Tdtm-Client/RM/stock"
)

func main() {
	settings := conf.New()
	order.Start(settings.OrderPort)
	stock.Start(settings.StockPort)
	account.Start(settings.AccountPort)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-c
}
