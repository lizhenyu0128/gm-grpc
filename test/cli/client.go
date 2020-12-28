package main

import (
	"context"
	"fmt"
	"github.com/Hyperledger-TWGC/grpc"
	"github.com/Hyperledger-TWGC/grpc/credentials"
	"github.com/Hyperledger-TWGC/grpc/test/hello"
	"time"
)

func main() {
	// single cert
	//creds, err := credentials.NewClientTLSFromFile("E:/gopath/projects/grpc/test/single-cert/ca.crt",
	//	"peer0.org3.example.com")

	// double cert
	creds, err := credentials.NewClientTLSFromFile("E:/gopath/projects/grpc/test/double-cert/server_ca.crt",
		"peer0.org1.example.com")
	if err != nil {
		fmt.Println("1", err)
		return
	}

	grpcOptions := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	//grpcOptions = append(grpcOptions, grpc.WithInsecure())

	ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
	conn, err := grpc.DialContext(ctx, "127.0.0.1:6262", grpcOptions...)
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}

	cli := hello.NewCommunicateClient(conn)
	ret, err := cli.Speak(context.Background(), &hello.Content{Detail: "aaa"})
	if err != nil {
		fmt.Println("speak error:", err)
		return
	}
	fmt.Println("speak ok:", ret.Detail)
}
