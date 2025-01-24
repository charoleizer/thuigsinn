package internal

import (
	"log"
	"net"

	"github.com/charoleizer/thuigsinn/ms-users/internal/handlers"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Infrastructure struct {
	Database   *mongo.Database
	Broker     nats.JetStreamContext
	ServerPort string
}

func RunServer(infrastructure Infrastructure) error {
	lis, err := net.Listen("tcp", ":"+infrastructure.ServerPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}

	s := grpc.NewServer()
	h := handlers.NewHandlers(infrastructure.Broker, infrastructure.Database)
	RegisterHandlers(s, h)
	reflection.Register(s)

	log.Printf("Server started at %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return err
	}

	return nil
}
