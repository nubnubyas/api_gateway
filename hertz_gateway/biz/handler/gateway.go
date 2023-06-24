package handler

import (
	"context"
	"fmt"

	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	// "github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
)

var SvcMap = make(map[string]genericclient.Client)

// input the service name and method name into nested map,
// gets the method name in the generic client
//var PathToMethod = make(map[string]map[string]string)

var pathToMethod = map[string]map[string]string{
	"student_api": {
		"query":  "queryStudent",
		"insert": "insertStudent",
	},
	"calculator": {
		"get": "calculate",
	},
}

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	// ie student api, calculator
	svcName := c.Param("svc")
	method := c.Param("method")
	fmt.Printf("%v\n", string(c.Request.Body()))
	methodName := pathToMethod[svcName][method]

	// get generic client
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, "cannot get generic client")
		return
	}

	// make generic call
	resp, err := cli.GenericCall(ctx, methodName, string(c.Request.Body()))
	if err != nil {
		fmt.Println("error here generic call")
		panic(err)
	}
	c.JSON(http.StatusOK, resp)
}
