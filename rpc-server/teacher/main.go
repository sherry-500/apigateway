package main

import (
	"log"
	"net"

	demo "github.com/Supernova114514/kitex_teacher/kitex_gen/demo/teacherservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	//"github.com/cloudwego/kitex/server/genericserver"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
   if err != nil {
	   log.Fatal(err)
   }

   addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")

   teacherService := &TeacherServiceImpl{}

   svr := demo.NewServer(teacherService, server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "teacherservice"}), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
