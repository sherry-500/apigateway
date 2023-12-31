// Code generated by hertz generator.

package demo

import (
	"context"
	"database/sql"
	"encoding/json"

	//"log"

	//"strings"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	db "github.com/sherry-500/apigateway/db/sqlc"

	//"github.com/cloudwego/kitex/pkg/generic"
	"github.com/sherry-500/apigateway/biz/client"
	//"github.com/sherry-500/apigateway/biz/idlMapping"
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

	//find the idl path
	_, err = db.StoreInstance.GetIdlmapByName(ctx, svcName)
	if err != nil {
		resp := &demo.ApiResp{
			Success: false,
			Message: "get idl failed:" + err.Error(),
		}
		if err == sql.ErrNoRows{
			c.JSON(consts.StatusNotFound, resp)
		}else{
			c.JSON(consts.StatusInternalServerError, resp)
		}
		return
	}

	client, err := client.GetClient(svcName)
	if err != nil {
		c.String(consts.StatusInternalServerError, "get client failed:" + err.Error())
		return
	}

	jsonData, err := json.Marshal(req)
	resp, err := client.GenericCall(ctx, methodName, string(jsonData))
	if err != nil {
		resp := &demo.ApiResp{
			Success: false,
			Message: "get resp faided:" + err.Error(),
		}
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	fmt.Println(resp)
	c.String(consts.StatusOK, resp.(string))
}
