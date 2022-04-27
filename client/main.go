//client（调用rpc的一方）
package main

import (
	"context"
	"doubna_userinfo/proto"
	"doubna_userinfo/tool"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

const (
	address = "localhost:50051"
)

var  Conn  *grpc.ClientConn
var UserClient   proto.UseCenterClient

func main() {
	//建立链接
	Conn,err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	UserClient= proto.NewUseCenterClient(Conn)
	defer Conn.Close()

	engine:=gin.Default()
	engine.POST("/api/user/login/pw",login)
	engine.Run(":8080")
}

func login(ctx *gin.Context) {

	loginAccount := ctx.PostForm("loginAccount")
	password := ctx.PostForm("password")


	if loginAccount == "" {
		tool.RespErrorWithData(ctx, "请输入注册时用的邮箱或者手机号")
		return
	}

	if password == "" {
		tool.RespErrorWithData(ctx, "请输入密码")
		return
	}


	loginResp,_:= UserClient.LoginByPW(context.Background(), &proto.LoginByPWReq{
	   LoginAccount: loginAccount,
	   PassWord: password,
   })


	if loginResp.OK {
		ctx.JSON(http.StatusOK, gin.H{
			"access_token":  "",
			"refresh_token": "",
			"token":         "",
			"status":        "ture",
			"data":          "",
		})
		return
	}else {
		ctx.JSON(http.StatusOK, gin.H{
			"access_token":  "",
			"refresh_token": "",
			"token":         "",
			"status":        "false",
			"data":          "",
		})
	}
}
