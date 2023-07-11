package main

import (
	"log"
	"net"

	calculator "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/calculator/calculatorservice"
	registerCenter "github.com/cloudwego/api_gateway/register_center/shared"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/cloudwego/kitex/server"
)

// rpc server 2 port:8082
func main() {

	if registerCenter.ErrRegistry != nil {
		klog.Fatal(registerCenter.ErrRegistry)
	}

	svr := calculator.NewServer(
		new(CalculatorApiImpl),
		server.WithRegistry(registerCenter.NacosRegistry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "calculator"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8082}),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
