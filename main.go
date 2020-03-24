package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"mj_logic_svr/points_mall"
	"net"
)

// 业务实现方法的容器
type server struct {
}

func (s *server) SayHello(ctx context.Context, req *points_mall.HelloRequest) (rsp *points_mall.HelloReply, err error) {
	rsp = &points_mall.HelloReply{}
	rsp.Message = "mj_logic" + req.Name
	err = nil
	fmt.Println("recv request ", req, rsp)
	return
}

func (s *server) SayHelloAgain(ctx context.Context, req *points_mall.HelloRequest) (rsp *points_mall.HelloReply, err error) {
	return
}

func main() {
	lis, err := net.Listen("tcp", ":8028") //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer() //创建gRPC服务

	/**注册接口服务
	 * 以定义proto时的service为单位注册，服务中可以有多个方法
	 * (proto编译时会为每个service生成Register***Server方法)
	 * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])
	 */
	points_mall.RegisterGreeterServer(s, &server{})
	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(s)
	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
