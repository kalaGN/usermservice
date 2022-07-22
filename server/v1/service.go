package main

import (
	"context"
	"fmt"
	"login"
	"net"

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
	l, err := net.Listen("tcp", ":8889")

	if err != nil {
		panic(err)
	}

	fmt.Println("listen on 127.0.0.1:8889")

	grpcServer := grpc.NewServer()

	loginService := LoginService{}
	login.RegisterLoginServiceServer(grpcServer, &loginService)

	err = grpcServer.Serve(l)

	if err != nil {
		println(err)
	}

}
