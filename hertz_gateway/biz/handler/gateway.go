package handler

import (
	"context"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
)

// type requiredParams struct {
// 	Method    string `form:"method,required" json:"method"`
// 	BizParams string `form:"biz_params,required" json:"biz_params"`
// }

var SvcMap = make(map[string]genericclient.Client)

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	// ie student api, calculator
	svcName := c.Param("svc")
	print(svcName + "\n")
	fmt.Printf("%v\n", c.Request.Body())
	// print("reached here\n")
	// if true {
	// 	c.JSON(http.StatusOK, "reached here")
	// 	return
	// }
	// print(c + "\n")
	// retrieve the correct client from SvcMap
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, "error gateway.go line 33")
		return
	}
	// var params requiredParams
	// if err := c.BindAndValidate(&params); err != nil {
	// 	hlog.Error(err)
	// 	c.JSON(http.StatusOK, ok)
	// 	return
	// }

	// Parse the raw bytes into a generic map
	var data map[string]interface{}
	err := json.Unmarshal(c.Request.Body(), &data)
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusOK, "error at line 52")
		return
	}

	// Read the request body into a byte slice
	// data, err := ioutil.ReadAll(c.Request.rawBody)
	// if err != nil {
	// 	// Handle the error
	// 	hlog.Error(err)
	// 	c.JSON(http.StatusOK, "error at line 60")
	// 	return
	// }

	resp, err := cli.GenericCall(ctx, "InsertStudent", data)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, resp)
}
