package middleware

import (
	"context"

	"github.com/cloudwego/api_gateway/hertz_gateway/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
)

// i think IdlMAP can be removed
var IdlMap = make(map[string]generic.DescriptorProvider)

func ProtocolTranslation() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// Parse IDL with Local Files
		// YOUR_IDL_PATH thrift file path, eg:./idl/example.thrift
		/*
			idlname := ctx.Param("svc")
			idlprovider := IdlMap[idlname]
		*/
		var req api.InsertStudentRequest
		err := ctx.BindAndValidate(&req)
		if err != nil {
			ctx.String(400, err.Error())
			return
		}

		// to comment out
		reqRpc := &api.InsertStudentRequest{
			Num:    req.Num,
			Name:   req.Name,
			Gender: req.Gender,
		}

		p, err := generic.NewThriftFileProvider("../idl/student_api.thrift")
		if err != nil {
			panic(err)
		}
		g, err := generic.JSONThriftGeneric(p)
		if err != nil {
			panic(err)
		}
		loadbalanceropt := client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer())
		cli, err := genericclient.NewClient("gateway_service", g, loadbalanceropt)
		if err != nil {
			panic(err)
		}
		// 'ExampleMethod' method name must be passed as param
		resp, err := cli.GenericCall(c, "", reqRpc)
		if err != nil {
			panic(err)
		}
		// ctx.Next()
		// resp is a JSON string
		ctx.JSON(200, resp)
	}
}
