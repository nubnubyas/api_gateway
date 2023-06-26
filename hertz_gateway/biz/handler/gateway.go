package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	// "github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
)

var SvcMap = make(map[string]genericclient.Client)

// input the service name and method name into nested map,
// gets the method name in the generic client
var PathToMethod = make(map[string]map[string]string)

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	// ie student api, calculator
	svcName := c.Param("svc")
	method := c.Param("method")
	fmt.Printf("%v\n", string(c.Request.Body()))
	path := svcName + "/" + method
	methodName := PathToMethod[svcName][path]

	// get generic client through service name
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, "cannot get generic client")
		return
	}

	// make generic call to the service with the method name
	resp, err := cli.GenericCall(ctx, methodName, string(c.Request.Body()))
	if err != nil {
		fmt.Println("error here generic call")
		panic(err)
	}
	c.JSON(http.StatusOK, resp)
}
