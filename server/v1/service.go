package main

import (
	"context"
	"encoding/json"
	"fmt"
	"login"
	"net"

	"pkg/common"
	"pkg/config"
	"pkg/database"

	"google.golang.org/grpc"
)

type LoginService struct {
}

func (s *LoginService) Login(ctx context.Context, req *login.Request) (*login.Response, error) {
	fmt.Printf("req name:%s, pwd:%s\n", req.Name, req.Pwd)

	accountInfo := database.GetUser(req.Name, req.Pwd)
	if accountInfo == nil {
		return &login.Response{
			Code:  "0",
			Msg:   "login fail,no users",
			Data:  "",
			Token: "",
		}, nil
	}
	b, err := json.Marshal(accountInfo)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &login.Response{
		Code:  "1",
		Msg:   "login success",
		Data:  string(b),
		Token: common.GenToken(),
	}, nil
}

func main() {

	port, _ := config.LoadPort()

	l, err := net.Listen("tcp", ":"+port)

	if err != nil {
		panic(err)
	}
	fmt.Println("listen on 127.0.0.1:" + port)

	grpcServer := grpc.NewServer()

	loginService := LoginService{}
	login.RegisterLoginServiceServer(grpcServer, &loginService)

	err = grpcServer.Serve(l)

	if err != nil {
		println(err)
	}

}
