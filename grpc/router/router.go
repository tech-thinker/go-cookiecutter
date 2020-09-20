package router

import (
	"log"
	"net"

	"github.com/mrasif/gomvc/grpc/todo"
	"github.com/mrasif/gomvc/service/initializer"
	"google.golang.org/grpc"
)

func Init(dependencies initializer.Services, port string) error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()

	t := todo.Init(dependencies)

	todo.RegisterTodoServiceServer(grpcServer, t)

	return grpcServer.Serve(listen)
}
