package main

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/sherry-500/kitex_student/kitex_gen/demo"
	"github.com/sherry-500/kitex_student/kitex_gen/demo/studentservice"
)

func main() {
	cli, err := studentservice.NewClient("student-server", client.WithHostPorts("127.0.0.1:8889"))
	if err != nil {
		panic("err init client:" + err.Error())
	}
	
	response, err := cli.Register(context.Background(), &demo.Student{
		Id: 1,
		Name: "aa",
		College: &demo.College{
			Name:    "computer school",
			Address: "computer school",
		},
		Email: []string{"aaaaaaaaa"},
		Sex: "hey",
	})
	if err != nil {
		panic("err query:" + err.Error())
	}
	klog.Infof("register: %v", response)
	resp, err := cli.Query(context.Background(), &demo.QueryReq{
		Id: 1,
	})
	if err != nil {
		panic("err query:" + err.Error())
	}
	klog.Infof("resp: %v", resp)
}
