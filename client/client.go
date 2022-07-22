package main

import (
	"context"
	"flag"
	"fmt"
	"login"
	"time"

	"google.golang.org/grpc"
)

//初始化测试数据
var (
	dest = flag.String("dest", ":8889", "The server address in the format of host:port")
	name = flag.String("name", "afei", "The name is login account name")
	pwd  = flag.String("pwd", "123456", "The passwd is the login account password to login")
	sign = flag.String("sign", "sign123123123", "The cmd is the login cmd")
)

func printLogin(client login.LoginServiceClient, req *login.Request) {
	fmt.Printf("Client: Request sign:%s, name:%s, pwd:%s\n", req.Sign, req.Name, req.Pwd)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.Login(ctx, req)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("res:" + response)
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*dest, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to Dial")
	}
	defer conn.Close()

	client := login.NewLoginServiceClient(conn)

	printLogin(client, &login.Request{
		Name: *name,
		Pwd:  *pwd,
		Sign: *sign,
	})

}
