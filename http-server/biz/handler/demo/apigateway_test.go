package demo

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route/param"
	"github.com/stretchr/testify/require"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	_ "github.com/lib/pq"
	db "github.com/sherry-500/apigateway/db/sqlc"
	"github.com/sherry-500/apigateway/util"
)

// 测试前需要将idl文件夹复制到handler/demo目录下

const URL = "/apigateway"

func TestApigateway(t *testing.T) {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	db.StoreInstance = db.NewStore(conn)
	
	testCases := []struct{
		name string
		reqBody string
		svc string
		method string
		checkResponse func(t *testing.T, c *app.RequestContext)
	}{
		{
			name: "OK",
			reqBody: `{"name":"sherry","id":1}`,
			svc: "studentservice",
			method: "Register",
			checkResponse: func(t *testing.T, c *app.RequestContext){
				require.Equal(t, consts.StatusOK, c.Response.StatusCode())
			},
		},
		{
			name: "InternalError",
			reqBody: `{"name":"sherry","id":1}`,
			svc: "studentservice",
			method: "Register",
			checkResponse: func(t *testing.T, c *app.RequestContext){
				require.Equal(t, consts.StatusInternalServerError, c.Response.StatusCode())
			},
		},
		{
			name: "NotFound",
			reqBody: `{"name":"sherry"}`,
			svc: "teacherservice",
			method: "Register",
			checkResponse: func(t *testing.T, c *app.RequestContext){
				require.Equal(t, consts.StatusNotFound, c.Response.StatusCode())
			},
		},
	}

	for i := range testCases {
		req := generateReq(testCases[i].reqBody, testCases[i].svc, testCases[i].method)
		Gateway(context.Background(), req)
		
		testCases[i].checkResponse(t, req)
	}

	//h.Close()
}

func generateReq(reqBody string, svcName string, methodName string) *app.RequestContext {
	req := &app.RequestContext{}

	req.SetFullPath(URL + "/" + svcName + "/" + methodName)
	req.Request.SetBody([]byte(reqBody))

	req.Request.Header.SetContentLength(len(reqBody))
	req.Request.SetHeader("Content-Type", "application/json")
	req.Request.Header.SetMethod(consts.MethodPost)
	req.Params = param.Params{
		{Key: "svcName", Value: svcName},
		{Key: "methodName", Value: methodName},
	}

	return req
}