package setup

import (
	"context"
	"log"
	"net"
	"os"

	route_handler "grpc-client/router"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"google.golang.org/grpc"
)

func Start() error {
	log := log.New(os.Stdout, "grpc skeleton : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	listen, err := net.Listen("tcp", viper.GetString("grpc.tcp_port"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting grpc server on port %s", viper.GetString("grpc.tcp_port"))

	srv := grpc.NewServer()
	defer srv.GracefulStop()

	mongodb, err := initMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	route_handler.RegisterGrpcRoute(srv, log, mongodb)

	return srv.Serve(listen)
}

func initMongoDB() (*mongo.Database, error) {
	ctx := context.Background()
	uri := viper.GetString("mongodb.uri")
	conn, err := mongo.Connect(ctx,
		options.Client().
			ApplyURI(uri).
			SetRetryReads(true).
			SetRetryWrites(true).
			SetWriteConcern(writeconcern.New(writeconcern.W(1))),
	)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := conn.Database(viper.GetString("mongodb.db"))
	return db, nil
}
