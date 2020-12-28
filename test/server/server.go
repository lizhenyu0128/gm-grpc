package main

import (
	"context"
	"fmt"
	"github.com/Hyperledger-TWGC/grpc"
	"github.com/Hyperledger-TWGC/grpc/credentials"
	"github.com/Hyperledger-TWGC/grpc/test/hello"
	"net"
)

type CommServer struct {
}

func (comm *CommServer)Speak(ctx context.Context, content *hello.Content) (*hello.Content, error) {
	fmt.Println("receive message :", content.Detail)
	return &hello.Content{Detail: "i am server"}, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6262")
	if err != nil {
		panic(err)
	}

	// load tls cert pair  single cert mode
	//creds, err := credentials.NewServerTLSFromFile("E:/gopath/projects/grpc/test/single-cert/server.crt",
	//	"E:/gopath/projects/grpc/test/single-cert/server.key")

	// load tls cert pair double cert mode
	creds, err := credentials.NewServerTLSFromFileDouble(
		"E:/gopath/projects/grpc/test/double-cert/server_sign.crt",
		"E:/gopath/projects/grpc/test/double-cert/server_sign.key",
		"E:/gopath/projects/grpc/test/double-cert/server_cipher.crt",
		"E:/gopath/projects/grpc/test/double-cert/server_cipher.key")
	if err != nil {
		panic(err)
	}

	grpcOptions := []grpc.ServerOption{grpc.Creds(creds)}
	gprcServer := grpc.NewServer(grpcOptions...)
	hello.RegisterCommunicateServer(gprcServer, &CommServer{})
	fmt.Println("beginning to serve ...")
	if err = gprcServer.Serve(l); err != nil {
		panic(err)
	}
}
