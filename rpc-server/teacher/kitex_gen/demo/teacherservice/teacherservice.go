// Code generated by Kitex v0.6.1. DO NOT EDIT.

package teacherservice

import (
	"context"
	demo "github.com/Supernova114514/kitex_teacher/kitex_gen/demo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return teacherServiceServiceInfo
}

var teacherServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "teacherService"
	handlerType := (*demo.TeacherService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register": kitex.NewMethodInfo(registerHandler, newTeacherServiceRegisterArgs, newTeacherServiceRegisterResult, false),
		"Query":    kitex.NewMethodInfo(queryHandler, newTeacherServiceQueryArgs, newTeacherServiceQueryResult, false),
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
	realArg := arg.(*demo.TeacherServiceRegisterArgs)
	realResult := result.(*demo.TeacherServiceRegisterResult)
	success, err := handler.(demo.TeacherService).Register(ctx, realArg.Teacher)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeacherServiceRegisterArgs() interface{} {
	return demo.NewTeacherServiceRegisterArgs()
}

func newTeacherServiceRegisterResult() interface{} {
	return demo.NewTeacherServiceRegisterResult()
}

func queryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.TeacherServiceQueryArgs)
	realResult := result.(*demo.TeacherServiceQueryResult)
	success, err := handler.(demo.TeacherService).Query(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeacherServiceQueryArgs() interface{} {
	return demo.NewTeacherServiceQueryArgs()
}

func newTeacherServiceQueryResult() interface{} {
	return demo.NewTeacherServiceQueryResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, teacher *demo.Teacher) (r *demo.RegisterResp, err error) {
	var _args demo.TeacherServiceRegisterArgs
	_args.Teacher = teacher
	var _result demo.TeacherServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Query(ctx context.Context, req *demo.QueryReq) (r *demo.Teacher, err error) {
	var _args demo.TeacherServiceQueryArgs
	_args.Req = req
	var _result demo.TeacherServiceQueryResult
	if err = p.c.Call(ctx, "Query", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
