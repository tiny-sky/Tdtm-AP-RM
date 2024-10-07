package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	ap "github.com/tiny-sky/Tdtm-Client/AP"
	"github.com/tiny-sky/Tdtm/client"
	"github.com/tiny-sky/Tdtm/core/registry"
	"github.com/tiny-sky/Tdtm/core/registry/etcdx"
)

func main() {
	serverUrl := flag.String("tdtm", "", "URL of the Easycar server")
	flag.Parse()

	var (
		d   registry.Discovery
		err error
	)

	d, err = etcdx.New(etcdx.Conf{
		Hosts: []string{"127.0.0.1:2379"}})
	if err != nil {
		return
	}

	ap.RegisterBuilder(d)

	cli, err := client.New(*serverUrl, client.WithDiscovery())
	if err != nil {
		return
	}

	defer func() {
		time.Sleep(3 * time.Minute)
		defer cli.Close(context.Background())
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	gid, err := cli.Begin(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Begin gid:", gid)

	if err = cli.Register(ctx, gid, ap.GetSrv()); err != nil {
		panic(err)
	}

	if err := cli.Start(ctx, gid); err != nil {
		fmt.Println("start err:", err)
	}
	fmt.Println("end gid:", gid)
	time.Sleep(3 * time.Second)
}
