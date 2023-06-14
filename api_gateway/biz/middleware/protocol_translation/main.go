package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

func ProtocolTranslation() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// Parse IDL with Local Files
		// YOUR_IDL_PATH thrift file path, eg:./idl/example.thrift
		p, err := generic.NewThriftFileProvider("../idl/student_api.thrift")
		if err != nil {
			panic(err)
		}
		g, err := generic.JSONThriftGeneric(p)
		if err != nil {
			panic(err)
		}
		cli, err := genericclient.NewClient("gateway_service", g)
		if err != nil {
			panic(err)
		}
		// 'ExampleMethod' method name must be passed as param
		resp, err := cli.GenericCall(context.Background(), "ExampleMethod", "{\"Msg\": \"hello\"}")
		// resp is a JSON string
		ctx.JSON(200, resp)
	}
}
