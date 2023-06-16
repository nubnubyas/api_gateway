package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
)

var IdlMap = make(map[string]generic.DescriptorProvider)

func ProtocolTranslation() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// Parse IDL with Local Files
		// YOUR_IDL_PATH thrift file path, eg:./idl/example.thrift
		/*
			idlname := ctx.Param("svc")
			idlprovider := IdlMap[idlname]
		*/
		p, err := generic.NewThriftFileProvider("../idl/student_api.thrift")
		if err != nil {
			panic(err)
		}
		g, err := generic.JSONThriftGeneric(p)
		if err != nil {
			panic(err)
		}
		opt := client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer())
		cli, err := genericclient.NewClient("gateway_service", g, opt)
		if err != nil {
			panic(err)
		}
		// 'ExampleMethod' method name must be passed as param
		resp, err := cli.GenericCall(c, "", "{\"Msg\": \"hello\"}")
		if err != nil {
			panic(err)
		}
		// ctx.Next()
		// resp is a JSON string
		ctx.JSON(200, resp)
	}
}
