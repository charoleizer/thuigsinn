package internal

import (
	"context"
	"users/cmd/users"
	"users/internal/handlers"

	"google.golang.org/grpc"
)

type server struct {
	users.UnimplementedUsersServer
	handlers *handlers.Handlers
}

func RegisterHandlers(s *grpc.Server, h *handlers.Handlers) {
	users.RegisterUsersServer(s, &server{handlers: h})
}

func (s *server) Create(ctx context.Context, req *users.CreateRequest) (*users.CreateResponse, error) {
	return s.handlers.Create(ctx, req)
}

func (s *server) Read(ctx context.Context, req *users.ReadRequest) (*users.ReadResponse, error) {
	return s.handlers.Read(ctx, req)
}

func (s *server) Update(ctx context.Context, req *users.UpdateRequest) (*users.UpdateResponse, error) {
	return s.handlers.Update(ctx, req)
}

func (s *server) Delete(ctx context.Context, req *users.DeleteRequest) (*users.DeleteResponse, error) {
	return s.handlers.Delete(ctx, req)
}
