package router

import (
	"log"
	"net"

	"github.com/tech-thinker/go-cookiecutter/grpc/todo"
	"github.com/tech-thinker/go-cookiecutter/service/initializer"
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
