package main

import (
	"net"
	"time"

	grader "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/grader/universitygrades"
	registerCenter "github.com/cloudwego/api_gateway/register_center/shared"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func main() {
	if registerCenter.ErrRegistry != nil {
		klog.Fatal(registerCenter.ErrRegistry)
	}

	svr := grader.NewServer(
		new(UniversityGradesImpl),
		server.WithRegistry(registerCenter.NacosRegistry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "grader"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8089}),
		server.WithReadWriteTimeout(600*time.Second),
	)

	err := svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
