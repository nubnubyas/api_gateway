package main

import (
	"net"
	"time"

	api "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/api/studentapi"
	registerCenter "github.com/cloudwego/api_gateway/register_center/shared"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/cloudwego/kitex/server"
)

func main() {
	if registerCenter.ErrRegistry != nil {
		klog.Fatal(registerCenter.ErrRegistry)
	}

	svr := api.NewServer(
		new(StudentApiImpl),
		server.WithRegistry(registerCenter.NacosRegistry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "student_api"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8081}),
		server.WithReadWriteTimeout(600*time.Second),
	)

	err := svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
