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

// var pathToMethod = make(map[string]string)

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	// ie student api, calculator
	svcName := c.Param("svc")
	//method := c.Param("method")
	fmt.Printf("%v\n", string(c.Request.Body()))
	// path := c.Request.URI().RequestURI()
	// methodName := pathToMethod[string(path)]

	// get generic client
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, "cannot get generic client")
		return
	}

	// make generic call
	resp, err := cli.GenericCall(ctx, "insertStudent", string(c.Request.Body()))
	if err != nil {
		fmt.Println("error here generic call")
		panic(err)
	}
	c.JSON(http.StatusOK, resp)
}
