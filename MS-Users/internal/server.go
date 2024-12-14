package internal

import (
	"log"
	"net"
	"users/internal/handlers"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(db *mongo.Database, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}

	s := grpc.NewServer()
	h := handlers.NewHandlers(db)
	RegisterHandlers(s, h)
	reflection.Register(s)

	log.Printf("Server started at %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return err
	}

	return nil
}
