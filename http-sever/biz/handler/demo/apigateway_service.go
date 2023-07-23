// Code generated by hertz generator.

package demo

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	demo "github.com/sherry-500/apigateway/biz/model/demo"
	"github.com/sherry-500/apigateway/biz/client"
)

// Gateway .
// @router /apigateway/:svcName/:methodName [POST]
func Gateway(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.ApiReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(demo.ApiResp)

	svcName := c.Param("svcName")
	methodName := c.Param("methodName")
	client.InitJsonGenericClient(svcName, "")
	kitexClient := client.Clients[svcName]

	kitexClient.GenericCall(ctx, methodName, req.Data)

	c.JSON(consts.StatusOK, resp)
}

// CreateIdl .
// @router /idl-manage/create [POST]
func CreateIdl(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.IdlMap
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(demo.IdlResp)
	fmt.Println("create")

	c.JSON(consts.StatusOK, resp)
}

// DeleteIdl .
// @router /idl-manage/delete [POST]
func DeleteIdl(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.Service
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(demo.IdlResp)

	c.JSON(consts.StatusOK, resp)
}

// UpdateIdl .
// @router /idl-manage/update [POST]
func UpdateIdl(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.IdlMap
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(demo.IdlResp)

	c.JSON(consts.StatusOK, resp)
}

// ResearchIdl .
// @router /idl-manage/research [POST]
func ResearchIdl(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.Service
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(demo.IdlMap)

	c.JSON(consts.StatusOK, resp)
}
