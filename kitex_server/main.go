package main

import (
	"log"
	"net"

	api "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/api/studentapi"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
)

func main() {
	r, err := registry.NewDefaultNacosRegistry()

	if err != nil {
		klog.Fatal(err)
	}

	svr := api.NewServer(
		&StudentApiImpl{},
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "studentManagement"}),
		server.WithServiceAddr(&net.TCPAddr{Port: 8081}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
