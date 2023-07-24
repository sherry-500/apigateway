package main

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
    "github.com/Supernova114514/kitex_teacher/kitex_gen/demo/teacherservice"
    "github.com/Supernova114514/kitex_teacher/kitex_gen/demo"


)

func main() {
	cli, err := teacherservice.NewClient("teacher-server", client.WithHostPorts("127.0.0.1:8890"))
	if err != nil {
		panic("err init client:" + err.Error())
	}

	resp, err := cli.Query(context.Background(), &demo.QueryReq{
		Id: 1,
	})
	if err != nil {
		panic("err query:" + err.Error())
	}
	klog.Infof("resp: %v", resp)
}
