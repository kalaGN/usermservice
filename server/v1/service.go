package main

import (
	"context"
	"fmt"
	"login"
	"net"
	"pkg/config"
	"pkg/database"

	"google.golang.org/grpc"
)

type LoginService struct {
}

func (s *LoginService) Login(ctx context.Context, req *login.Request) (*login.Response, error) {
	fmt.Printf("req name:%s, pwd:%s\n", req.Name, req.Pwd)

	return &login.Response{
		Code:  "1",
		Msg:   "login success",
		Data:  "ok",
		Token: "ttt",
	}, nil
}

func main() {

	database.GetUser("afei", "e10adc3949ba59abbe56e057f20f883e")

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
