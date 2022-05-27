package chat

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"grpc-client/domain"
	"grpc-client/domain/repo"
	"grpc-client/pb/chat"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatHandler struct {
	ctx  context.Context
	log  *log.Logger
	repo domain.MongoSessionRepo
}

func NewChatHandler(log *log.Logger, mongodb *mongo.Database) *ChatHandler {
	repo := repo.NewMonogoSessionRepo(mongodb, log)
	return &ChatHandler{ctx: context.Background(), log: log, repo: repo}
}

func (c *ChatHandler) ReplyMessage(ctx context.Context, in *chat.HelloRequest) (*chat.HelloResponse, error) {
	c.log.Println("ReplyMessage:", in.Name)
	return &chat.HelloResponse{Message: "Hello " + in.Name}, nil
}

func (c *ChatHandler) CheckMyData(ctx context.Context, in *chat.CheckMyDataRequest) (*chat.CheckMyDataResponse, error) {
	c.log.Println("CheckMyData:", in)

	ss, err := c.repo.CreateSession()
	if err != nil {
		return nil, err
	}

	res := &chat.CheckMyDataResponse{}
	err = mongo.WithSession(c.ctx, ss, func(smc mongo.SessionContext) error {
		j, _ := json.Marshal(in)
		var bsData interface{}
		_ = bson.UnmarshalExtJSON(j, true, &bsData)
		err := c.repo.InsertLogs(smc, bson.M{"name": "CheckMyData", "data": bsData, "date": time.Now()})
		if err != nil {
			c.log.Fatal(err)
		}

		if in.Age > 16 {
			res.Code = 200
			res.Message = "You are allowed to enter"
			if err = ss.CommitTransaction(smc); err != nil {
				return err
			}
		} else {
			if err = ss.AbortTransaction(smc); err != nil {
				return err
			}

			res.Code = 403
			res.Message = "You are not allowed to enter"
		}

		return nil
	})
	if err != nil {
		c.log.Fatal(err)
	}
	ss.EndSession(ctx)
	return res, nil
}
