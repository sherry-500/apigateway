// Code generated by hertz generator. DO NOT EDIT.

package demo

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	demo "github.com/sherry-500/apigateway/biz/handler/demo"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_apigateway := root.Group("/apigateway", _apigatewayMw()...)
		{
			_svcname := _apigateway.Group("/:svcName", _svcnameMw()...)
			_svcname.POST("/:methodName", append(_gatewayMw(), demo.Gateway)...)
		}
	}
}