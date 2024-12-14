package repositories

import (
	"context"
	"users/cmd/users"
	"users/internal/dtos"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersRepository interface {
	Create(ctx context.Context, users dtos.ExtendedCreateRequest) (primitive.ObjectID, error)
	Read(ctx context.Context, id primitive.ObjectID) (*users.ReadResponse, error)
	Update(ctx context.Context, id primitive.ObjectID, users dtos.ExtendedUpdateRequest) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
