package handler

import (
	"context"
	"fmt"

	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
)

type requiredParams struct {
	Method    string `form:"method,required" json:"method"`
	BizParams string `form:"biz_params,required" json:"biz_params"`
}

var SvcMap = make(map[string]genericclient.Client)

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	// ie student api, calculator
	svcName := c.Param("svc")
	//method := c.Param("method")
	fmt.Printf("%v\n", string(c.Request.Body()))

	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, "error gateway.go line 33")
		return
	}

	var params requiredParams
	if err := c.BindAndValidate(&params); err != nil {
		hlog.Error(err)
		c.JSON(http.StatusOK, ok)
		return
	}
	fmt.Println("binded")

	resp, err := cli.GenericCall(ctx, "insertStudent", string(c.Request.Body()))
	if err != nil {
		fmt.Println("error here generic call")
		panic(err)
	}
	c.JSON(http.StatusOK, resp)
}
