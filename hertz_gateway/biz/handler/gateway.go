package handler

import (
	"context"

	"encoding/json"
	"net/http"

	errors "github.com/cloudwego/api_gateway/error"
	"github.com/cloudwego/api_gateway/error/kitex_gen/common"
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
)

// SvcMap is a map of service name to generic client
var SvcMap = make(map[string]genericclient.Client)

// input the service name and method name into nested map,
// gets the method name in the generic client
// var PathToMethod = make(map[string]map[string]string)
var PathToMethod = make(map[string]map[MethodPath]string)

var FileToSvc = make(map[string][]string)

type MethodPath struct {
	Path   string
	Method string
}

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {

	reqBody := string(c.Request.Body())

	var methodParam MethodPath

	// Get the HTTP request path (works)
	httpPath := string(c.Request.URI().RequestURI())
	//fmt.Println(string(pathtest))
	// Get the HTTP request method (works)
	httpMethod := string(c.Request.Method())
	//fmt.Println(string(methodtest()))
	methodParam.Method = httpMethod
	methodParam.Path = httpPath

	// verify the if request body is encoded in JSON (only if it is non-GET requests)
	// GET requests do not have request body (uses query string instead)
	if !checkJSON(reqBody) {
		hlog.Error("JsonNotFound err")
		c.JSON(http.StatusNoContent, errors.New(common.Err_JsonNotFound))
		return
	}

	// ie student_api, calculators
	svcName := c.Param("svc")
	// ie queryStudent, insertStudent, get
	// method := c.Param("method")
	// path := svcName + "/" + method
	methodName := PathToMethod[svcName][methodParam]

	// get generic client through service name
	cli, ok := SvcMap[svcName]
	if !ok {
		hlog.Errorf("Generic Client Not Found err")
		c.JSON(http.StatusNotFound, errors.New(common.Err_GenericClientNotFound))
		return
	}

	// make generic call to the service with the method name
	resp, err := cli.GenericCall(ctx, methodName, reqBody)
	if err != nil {
		hlog.Errorf("Generic Call err:%v", err)
		c.JSON(http.StatusBadRequest, errors.New(common.Err_GenericCallFailed))
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
