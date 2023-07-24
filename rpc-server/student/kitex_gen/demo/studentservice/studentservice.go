// Code generated by Kitex v0.6.1. DO NOT EDIT.

package studentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	demo "github.com/sherry-500/kitex_student/kitex_gen/demo"
)

func serviceInfo() *kitex.ServiceInfo {
	return studentServiceServiceInfo
}

var studentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "StudentService"
	handlerType := (*demo.StudentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register": kitex.NewMethodInfo(registerHandler, newStudentServiceRegisterArgs, newStudentServiceRegisterResult, false),
		"Query":    kitex.NewMethodInfo(queryHandler, newStudentServiceQueryArgs, newStudentServiceQueryResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "demo",
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

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.StudentServiceRegisterArgs)
	realResult := result.(*demo.StudentServiceRegisterResult)
	success, err := handler.(demo.StudentService).Register(ctx, realArg.Student)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceRegisterArgs() interface{} {
	return demo.NewStudentServiceRegisterArgs()
}

func newStudentServiceRegisterResult() interface{} {
	return demo.NewStudentServiceRegisterResult()
}

func queryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.StudentServiceQueryArgs)
	realResult := result.(*demo.StudentServiceQueryResult)
	success, err := handler.(demo.StudentService).Query(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceQueryArgs() interface{} {
	return demo.NewStudentServiceQueryArgs()
}

func newStudentServiceQueryResult() interface{} {
	return demo.NewStudentServiceQueryResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, student *demo.Student) (r *demo.RegisterResp, err error) {
	var _args demo.StudentServiceRegisterArgs
	_args.Student = student
	var _result demo.StudentServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Query(ctx context.Context, req *demo.QueryReq) (r *demo.Student, err error) {
	var _args demo.StudentServiceQueryArgs
	_args.Req = req
	var _result demo.StudentServiceQueryResult
	if err = p.c.Call(ctx, "Query", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
