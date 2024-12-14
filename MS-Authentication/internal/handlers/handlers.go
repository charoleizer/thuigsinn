package handlers

import (
	"context"
	"time"

	"github.com/charoleizer/thuigsinn/ms-authentication/internal/repositories"
	"github.com/charoleizer/thuigsinn/ms-authentication/pkg/proto/authentication"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CreateTimeout = time.Second * 5
)

type Handlers struct {
	db *mongo.Database
}

func NewHandlers(db *mongo.Database) *Handlers {
	return &Handlers{
		db: db,
	}
}

func (h *Handlers) Register(req *authentication.RegisterRequest, resp *authentication.RegisterResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), CreateTimeout)
	defer cancel()

	usersRepository := repositories.NewUsers(h.db)
	err := usersRepository.Create(ctx, req)
	if err != nil {
		resp.Status = authentication.Status_Failed
		return err
	}

	resp.Status = authentication.Status_Completed
	return nil
}
