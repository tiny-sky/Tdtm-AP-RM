package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	ap "github.com/tiny-sky/Tdtm-Client/AP"
	"github.com/tiny-sky/Tdtm/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// go run dircet.go --tdtm="127.0.0.1:8088"
	serverUrl := flag.String("tdtm", "", "URL of the Easycar server")
	flag.Parse()

	var opts []client.Option
	opts = append(opts, client.WithGrpcDailOpts([]grpc.DialOption{grpc.WithBlock(), grpc.WithReturnConnectionError()}))
	opts = append(opts, client.WithGrpcDailOpts([]grpc.DialOption{grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)}))
	opts = append(opts, client.WithGrpcDailOpts([]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}))
	opts = append(opts, client.WithConnTimeout(5*time.Second))

	cli, err := client.New(*serverUrl, opts...)
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}

	ctx := context.Background()
	defer cli.Close(ctx)

	gid, err := cli.Begin(ctx)
	if err != nil {
		log.Fatal("Failed to begin transaction:", err)
	}
	fmt.Println("Begin gid:", gid)

	// 注册服务到事务中
	if err = cli.Register(ctx, gid, ap.GetSrv()); err != nil {
		log.Fatal("Failed to register service:", err)
	}

	// 启动事务
	if err := cli.Start(ctx, gid); err != nil {
		fmt.Println("Start error:", err)
	}

	// 打印结束 gid
	fmt.Println("End gid:", gid)
}
