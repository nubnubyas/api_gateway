package main

import (
	"log"
	"net"

	api "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/api/studentapi"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/registry-nacos/registry"

	"github.com/cloudwego/kitex/server"
)

// rpc server 1
func main() {

	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		klog.Fatal(err)
	}

	svr := api.NewServer(
		new(StudentApiImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "student_api"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8081}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
