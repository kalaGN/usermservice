/*
 * @Author: afei
 * @Date: 2022-07-22 10:37:52
 * @LastEditors: afei
 * @LastEditTime: 2022-07-29 17:20:56
 * @Description:
 *
 * Copyright (c) 2022 , All Rights Reserved.
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"login"
	"pkg/config"
	"time"

	"google.golang.org/grpc"
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

	fmt.Println(response)
}

func main() {
	port, _ := config.LoadPort()
	flag.Parse()
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to Dial")
	}
	defer conn.Close()

	client := login.NewLoginServiceClient(conn)

	printLogin(client, &login.Request{
		Name: "afei",
		Pwd:  "pwd",
		Sign: "sign123",
	})

}
