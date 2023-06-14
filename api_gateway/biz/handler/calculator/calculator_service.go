// Code generated by hertz generator.

package calculator

import (
	"context"

	calculator "github.com/cloudwego/api_gateway/api_gateway/biz/model/calculator"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Calculate .
// @router calculator/get [GET]
func Calculate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req calculator.CalculatorReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(calculator.CalculatorResp)

	c.JSON(consts.StatusOK, resp)
}
