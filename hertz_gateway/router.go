// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cloudwego/api_gateway/hertz_gateway/biz/handler"
	registerCenter "github.com/cloudwego/api_gateway/register_center/shared"
	"github.com/cloudwego/thriftgo/parser"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/", func(ctx context.Context, c *app.RequestContext) {
		registerIDLs(r)
		c.JSON(http.StatusOK, "api-gateway is updated and running ... ")
	})

	print("customizedRegister\n")
	registerGateway(r)
}

// to update the IDL mapping
func registerIDLs(r *server.Hertz) {
	// if handler.SvcMap == nil {
	handler.SvcMap = make(map[string]genericclient.Client)
	// }

	// if handler.PathToMethod == nil {
	handler.PathToMethod = make(map[string]map[string]string)
	// }

	idlPath := "../idl/"
	c, err := os.ReadDir(idlPath)
	if err != nil {
		hlog.Fatalf("new thrift file provider failed: %v", err)
	}

	// same resolver for all generic clients
	if registerCenter.ErrResolver != nil {
		hlog.Fatalf("err:%v", err)
	}

	// generic clients creation
	for _, entry := range c {

		svcName := strings.ReplaceAll(entry.Name(), ".thrift", "")

		filePath := idlPath + entry.Name()

		fileSyntax, err := parser.ParseFile(filePath, nil, false)
		if err != nil {
			hlog.Fatalf("parse file failed: %v", err)
			break
		}

		fileSyntax.ForEachService(func(v *parser.Service) bool {
			v.ForEachFunction(func(v *parser.Function) bool {
				functionName := v.Name
				if handler.PathToMethod[svcName] == nil {
					handler.PathToMethod[svcName] = make(map[string]string)
				}

				switch {
				case len(v.Annotations.Get("api.get")) > 0:
					Subpath := methodSplit(v.Annotations.Get("api.get"))
					handler.PathToMethod[svcName][Subpath] = functionName
				case len(v.Annotations.Get("api.post")) > 0:
					Subpath := methodSplit(v.Annotations.Get("api.post"))
					handler.PathToMethod[svcName][Subpath] = functionName
				default:
					// Use a default HTTP method type
				}
				return true
			})
			return true
		})

		//print out the mapping
		fmt.Println(handler.PathToMethod)

		provider, err := generic.NewThriftFileProvider(entry.Name(), idlPath)
		if err != nil {
			hlog.Fatalf("new thrift file provider failed: %v", err)
			break
		}

		g, err := generic.JSONThriftGeneric(provider)
		if err != nil {
			hlog.Fatal(err)
		}

		loadbalanceropt := client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer())
		// creates new generic client for each IDL
		cli, err := genericclient.NewClient(
			svcName,
			g,
			client.WithResolver(registerCenter.NacosResolver),
			loadbalanceropt,
		)
		if err != nil {
			hlog.Fatal(err)
		}

		handler.SvcMap[svcName] = cli
		fmt.Println(svcName)
	}
}

// to register and establish routing for gateway
func registerGateway(r *server.Hertz) {
	group := r.Group("/")
	{
		group.Any("/:svc/:method", handler.Gateway)
	}

	registerIDLs(r)

	print("registered gateway\n")
}

func methodSplit(pathName []string) string {
	path := pathName[0]
	parts := strings.Split(path, "/")
	subpath := strings.Join(parts[1:], "/")
	return subpath
}
