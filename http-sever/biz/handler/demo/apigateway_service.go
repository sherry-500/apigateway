// Code generated by hertz generator.

package demo

import (
	"context"
	"encoding/json"
	//"fmt"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sherry-500/apigateway/biz/client"
	"github.com/sherry-500/apigateway/biz/idlMapping"
	demo "github.com/sherry-500/apigateway/biz/model/demo"
)

// Gateway .
// @router /apigateway/:svcName/:methodName [POST]
func Gateway(ctx context.Context, c *app.RequestContext) {
	var err error
	//var req demo.ApiReq
	var req interface{}
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	svcName := c.Param("svcName")
	methodName := c.Param("methodName")

	if idlPath, ok := idlMapping.IDLMap[svcName]; ok{
		client.InitJsonGenericClient(svcName, idlPath)
	}else{
		log.Fatal("there is no service")
	}

	kitexClient, ok := client.Clients[svcName]
	if !ok {
		log.Fatal("no client")
	}

	//fmt.Println(req.Data)
	jsonData, err := json.Marshal(req)
	resp, err := kitexClient.GenericCall(ctx, methodName, string(jsonData))
	if err != nil {
		panic("err rpc server:" + err.Error())
	}

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
	svcName := req.SvcName
	idlPath := req.Path
	ok := idlMapping.AddIdl(svcName, idlPath)
	if !ok {
		c.JSON(consts.StatusInternalServerError, resp)
	}
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
	svcName := req.Name
	ok := idlMapping.DelIdl(svcName)
	if !ok {
		c.JSON(consts.StatusInternalServerError, resp)
	}
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
	svcName := req.SvcName
	idlPath := req.Path
	ok := idlMapping.UpdateIdl(svcName, idlPath)
	if !ok {
		c.JSON(consts.StatusInternalServerError, resp)
	}
	c.JSON(consts.StatusOK, resp)
}

// SearchIdl .
// @router /idl-manage/search [POST]
func SearchIdl(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.Service
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(demo.IdlMap)
	svcName := req.Name
	idlPath := idlMapping.GetIdl(svcName)
	
	resp.SvcName = svcName
	resp.Path = idlPath
	c.JSON(consts.StatusOK, resp)
}
