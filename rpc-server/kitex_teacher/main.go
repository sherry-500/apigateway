package main

import (
	// demo "github.com/Supernova114514/kitex_teacher/kitex_gen/demo/teacherservice"
	"log"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/cloudwego/kitex/server"
	"net"
   "github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/generic"
    "github.com/cloudwego/kitex/server/genericserver"
)

func main() {
	//svr := demo.NewServer(new(TeacherServiceImpl))

	// 本地文件 idl 解析
   // YOUR_IDL_PATH thrift 文件路径: e.g. ./idl/example.thrift
   p, err := generic.NewThriftFileProvider("../../idl/kitex_teacher.thrift")
   if err != nil {
	   panic(err)
   }
   // 构造 JSON 请求和返回类型的泛化调用
   g, err := generic.JSONThriftGeneric(p)
   if err != nil {
	   panic(err)
   }

   r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
   if err != nil {
	   log.Fatal(err)
   }

   addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")

   teacherService := &TeacherServiceImpl{}
   teacherService.InitDB()

   //svr := demo.NewServer(teacherService, server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "teacherservice"}), server.WithServiceAddr(addr))
   svr := genericserver.NewServer(teacherService, g, server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "teacherservice"}), server.WithServiceAddr(addr))
   //  svr := demo.NewServer(new(TeacherServiceImpl), server.WithServiceAddr(addr))

   err = svr.Run()

   if err != nil {
	   log.Println(err.Error())
   }
}
