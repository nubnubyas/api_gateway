// Code generated by Kitex v0.6.1. DO NOT EDIT.

package studentapi

import (
	"context"
	api "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return studentApiServiceInfo
}

var studentApiServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "StudentApi"
	handlerType := (*api.StudentApi)(nil)
	methods := map[string]kitex.MethodInfo{
		"queryStudent":  kitex.NewMethodInfo(queryStudentHandler, newStudentApiQueryStudentArgs, newStudentApiQueryStudentResult, false),
		"insertStudent": kitex.NewMethodInfo(insertStudentHandler, newStudentApiInsertStudentArgs, newStudentApiInsertStudentResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func queryStudentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.StudentApiQueryStudentArgs)
	realResult := result.(*api.StudentApiQueryStudentResult)
	success, err := handler.(api.StudentApi).QueryStudent(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentApiQueryStudentArgs() interface{} {
	return api.NewStudentApiQueryStudentArgs()
}

func newStudentApiQueryStudentResult() interface{} {
	return api.NewStudentApiQueryStudentResult()
}

func insertStudentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.StudentApiInsertStudentArgs)
	realResult := result.(*api.StudentApiInsertStudentResult)
	success, err := handler.(api.StudentApi).InsertStudent(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentApiInsertStudentArgs() interface{} {
	return api.NewStudentApiInsertStudentArgs()
}

func newStudentApiInsertStudentResult() interface{} {
	return api.NewStudentApiInsertStudentResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) QueryStudent(ctx context.Context, req *api.QueryStudentRequest) (r *api.QueryStudentResponse, err error) {
	var _args api.StudentApiQueryStudentArgs
	_args.Req = req
	var _result api.StudentApiQueryStudentResult
	if err = p.c.Call(ctx, "queryStudent", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) InsertStudent(ctx context.Context, req *api.InsertStudentRequest) (r *api.InsertStudentResponse, err error) {
	var _args api.StudentApiInsertStudentArgs
	_args.Req = req
	var _result api.StudentApiInsertStudentResult
	if err = p.c.Call(ctx, "insertStudent", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
