package main

import (
	"log"
	"net"

	calculator "github.com/cloudwego/api_gateway/kitex_server/kitex_server_2/kitex_gen/calculator/calculatorservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/registry-nacos/registry"

	"github.com/cloudwego/kitex/server"
)

// rpc server 2 port:8082
func main() {

	// might need to move this out (under kitex_server)
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		klog.Fatal(err)
	}

	svr := calculator.NewServer(
		new(CalculatorApiImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "calculator"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8082}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
