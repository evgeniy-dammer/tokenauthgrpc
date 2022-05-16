package main

import (
	"fmt"
	"log"
	"net"

	"github.com/evgeniy-dammer/tokenauthgrpc/productservice/handlers"
	"github.com/evgeniy-dammer/tokenauthgrpc/productservice/interceptors"
	productservice "github.com/evgeniy-dammer/tokenauthgrpc/productservice/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("../../keys/server-cert.pem", "../../keys/server-key.pem")

	if err != nil {
		log.Fatalf("Failed to setup tls: %v", err)
	}

	listen, err := net.Listen("tcp", ":1111")

	if err != nil {
		fmt.Println(err)
	}

	defer listen.Close()

	productServ := handlers.ProductServiceServer{}

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptors.TokenAuthInterceptor,
			),
		),
	)

	productservice.RegisterProductServiceServer(grpcServer, &productServ)

	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
