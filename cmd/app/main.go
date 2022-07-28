package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	protoUserService "github.com/theartofdevel/grpc-contracts/gen/go/user_service/service/v1"
	v1 "github.com/theartofdevel/grpc-service/internal/controller/grpc/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcHostPort = "0.0.0.0:8082"
)

func main() {
	grpcServer := grpc.NewServer()
	listen, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		panic(err)
	}

	protoUserService.RegisterUserServiceServer(
		grpcServer,
		v1.NewUserServer(protoUserService.UnimplementedUserServiceServer{}),
	)

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = protoUserService.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, grpcHostPort, opts)
	if err != nil {
		panic(err)
	}

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		return http.ListenAndServe(":8081", mux)
	})

	err = g.Wait()
	if err != nil {
		panic(err)
	}
}
