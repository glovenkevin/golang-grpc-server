package chat

import (
	"context"
	"log"

	"grpc-client/pb/chat"
)

type ChatHandler struct {
	log *log.Logger
}

func NewChatHandler(log *log.Logger) *ChatHandler {
	return &ChatHandler{log: log}
}

func (c *ChatHandler) ReplyMessage(ctx context.Context, in *chat.HelloRequest) (*chat.HelloResponse, error) {
	c.log.Println("ReplyMessage:", in.Name)
	return &chat.HelloResponse{Message: "Hello " + in.Name}, nil
}

func (c *ChatHandler) CheckMyData(ctx context.Context, in *chat.CheckMyDataRequest) (*chat.CheckMyDataResponse, error) {
	c.log.Println("CheckMyData:", in)
	res := &chat.CheckMyDataResponse{}
	if in.Age > 16 {
		res.Code = 200
		res.Message = "You are allowed to enter"
	} else {
		res.Code = 403
		res.Message = "You are not allowed to enter"
	}
	return res, nil
}
