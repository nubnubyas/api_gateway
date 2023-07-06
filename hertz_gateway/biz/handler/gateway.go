package handler

import (
	"context"

	"encoding/json"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"

	// "github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
)

// SvcMap is a map of service name to generic client
var SvcMap = make(map[string]genericclient.Client)

// input the service name and method name into nested map,
// gets the method name in the generic client
var PathToMethod = make(map[string]map[string]string)

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {

	reqBody := string(c.Request.Body())

	// verify the if request body is encoded in JSON (only if it is non-GET requests)
	// GET requests do not have request body (uses query string instead)
	if string(c.Request.Method()) != "GET" {
		if !checkJSON(reqBody) {
			c.JSON(http.StatusOK, "request body is not in JSON format")
			return
		}
	}

	// ie student_api, calculator
	svcName := c.Param("svc")
	// ie queryStudent, insertStudent, get
	method := c.Param("method")
	path := svcName + "/" + method
	methodName := PathToMethod[svcName][path]

	// get generic client through service name
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, "cannot get generic client")
		return
	}

	// make generic call to the service with the method name
	resp, err := cli.GenericCall(ctx, methodName, reqBody)
	if err != nil {
		c.JSON(http.StatusOK, "error here generic call")
		panic(err)
	}

	c.JSON(http.StatusOK, resp)
}

// return true if the string (req body) is in JSON format
func checkJSON(data string) bool {
	var temp interface{}
	err := json.Unmarshal([]byte(data), &temp)
	return err == nil
}
