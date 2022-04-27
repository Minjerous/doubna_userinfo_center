
package main

import (
	"context"
	"doubna_userinfo/proto"
	"doubna_userinfo/server"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	address = "localhost:50051"
)
func main() {
	// 监听端口
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() //获取新服务示例
	proto.RegisterUseCenterServer(s,&UserServer{})
	// 开始处理
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type UserServer struct {
	proto.UnimplementedUseCenterServer
}



func  (s *UserServer) LoginByPW(ctx context.Context,req *proto.LoginByPWReq)(*proto.LoginByPWResp, error){

	resp := &proto.LoginByPWResp{}

	_, flag, err := server.JudgePasswordCorrect(req.LoginAccount, req.PassWord)

	if err != nil {
		resp.OK=false
		fmt.Println("UserLogin is", err)
		return  resp,err
	}

	if flag {
		resp.OK=true
		return  resp,err
	}else {
		resp.OK=false
		return  resp,err
	}
}
