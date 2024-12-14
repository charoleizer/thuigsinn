package repositories

import (
	"context"

	"github.com/charoleizer/thuigsinn/ms-authentication/internal/errdefs"
	"github.com/charoleizer/thuigsinn/ms-authentication/internal/models"
	"github.com/charoleizer/thuigsinn/ms-authentication/pkg/proto/authentication"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersRepository struct {
	db *mongo.Database
}

func NewUsers(db *mongo.Database) UsersRepository {
	return &usersRepository{db}
}

func (u *usersRepository) Create(ctx context.Context, users *authentication.RegisterRequest) error {
	objectID, err := primitive.ObjectIDFromHex(users.Userid)
	if err != nil {
		return errdefs.ErrInvalidObjectID(err.Error())
	}

	model := models.User{
		ID:       objectID,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}

	collection := u.db.Collection(model.TableName())

	// Simulate a lazy database
	// time.Sleep(3 * time.Second)

	_, err = collection.InsertOne(ctx, model)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return errdefs.ErrDatabaseTimeout("collection.InsertOne")
		}
		return errdefs.ErrDatabaseInsertOne(err.Error())
	}

	return nil
}
