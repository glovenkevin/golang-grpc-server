package main

import (
	"log"
	"net"
	"os"

	chat_handler "grpc-client/chat"
	chat "grpc-client/pb/chat"

	"google.golang.org/grpc"
)

func main() {
	log := log.New(os.Stdout, "grpc skeleton : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	log.Println("Starting server...")

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	grpcRoute(srv, log)

	defer srv.GracefulStop()

	if err = srv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
}

func grpcRoute(srv *grpc.Server, log *log.Logger) {
	handler := chat_handler.NewChatHandler(log)
	chat.RegisterChatServiceServer(srv, handler)
}
