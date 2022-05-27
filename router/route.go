package router

import (
	chat_handler "grpc-client/chat"
	"grpc-client/pb/chat"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func RegisterGrpcRoute(srv *grpc.Server, log *log.Logger, mongodb *mongo.Database) {
	handler := chat_handler.NewChatHandler(log, mongodb)
	chat.RegisterChatServiceServer(srv, handler)
}
