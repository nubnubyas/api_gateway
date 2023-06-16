package handler

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

type requiredParams struct {
	Method    string `form:"method,required" json:"method"`
	BizParams string `form:"biz_params,required" json:"biz_params"`
}

var SvcMap = make(map[string]genericclient.Client)

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	svcName := c.Param("svc")
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, ok)
		return
	}
	var params requiredParams
	if err := c.BindAndValidate(&params); err != nil {
		hlog.Error(err)
		c.JSON(http.StatusOK, ok)
		return
	}

	req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(params.BizParams)))
	if err != nil {
		hlog.Warnf("new http request failed: %v", err)
		c.JSON(http.StatusOK, ok)
		return
	}

	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		hlog.Errorf("convert request failed: %v", err)
		c.JSON(http.StatusOK, ok)
		return
	}
	resp, err := cli.GenericCall(ctx, "", customReq)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, resp)
}
