package repositories

import (
	"context"

	"github.com/charoleizer/thuigsinn/ms-users/internal/dtos"
	"github.com/charoleizer/thuigsinn/ms-users/internal/errdefs"
	"github.com/charoleizer/thuigsinn/ms-users/internal/models"
	"github.com/charoleizer/thuigsinn/ms-users/pkg/proto/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersRepository struct {
	db *mongo.Database
}

func NewUsers(db *mongo.Database) UsersRepository {
	return &usersRepository{db}
}

func (u *usersRepository) _isEmailAlreadyExists(model models.User) (bool, error) {
	collection := u.db.Collection(model.TableName())

	filter := bson.M{"email": model.Email}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (u *usersRepository) _isUsernameAlreadyExists(model models.User) (bool, error) {
	collection := u.db.Collection(model.TableName())

	filter := bson.M{"username": model.Username}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (u *usersRepository) Create(ctx context.Context, users dtos.ExtendedCreateRequest) (primitive.ObjectID, error) {
	model := models.User{
		ID:       primitive.NewObjectID(),
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}

	exists, err := u._isEmailAlreadyExists(model)
	if err != nil {
		return primitive.NilObjectID, errdefs.ErrDatabaseCountDocuments(err.Error())
	}
	if exists {
		return primitive.NilObjectID, errdefs.ErrEmailAlreadyExists
	}

	exists, err = u._isUsernameAlreadyExists(model)
	if err != nil {
		return primitive.NilObjectID, errdefs.ErrDatabaseCountDocuments(err.Error())
	}
	if exists {
		return primitive.NilObjectID, errdefs.ErrUsernameAlreadyExists
	}

	collection := u.db.Collection(model.TableName())

	result, err := collection.InsertOne(ctx, model)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return primitive.NilObjectID, errdefs.ErrDatabaseTimeout("collection.InsertOne")
		}
		return primitive.NilObjectID, errdefs.ErrDatabaseInsertOne(err.Error())
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (u *usersRepository) Read(ctx context.Context, id primitive.ObjectID) (*users.ReadResponse, error) {
	var model *models.User

	collection := u.db.Collection(model.TableName())

	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errdefs.ErrUserNotFound
		}
		return nil, errdefs.ErrDatabaseFindOne(err.Error())
	}

	return &users.ReadResponse{
		Username: model.Username,
		Email:    model.Email,
	}, nil

}

func (u *usersRepository) Update(ctx context.Context, id primitive.ObjectID, users dtos.ExtendedUpdateRequest) error {
	var model *models.User

	collection := u.db.Collection(model.TableName())

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{}}

	if users.Username != "" {
		update["$set"].(bson.M)["username"] = users.Username
	}

	if users.Email != "" {
		update["$set"].(bson.M)["email"] = users.Email
	}

	if len(update["$set"].(bson.M)) == 0 {
		return nil
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errdefs.ErrUserNotFound
		}
		return errdefs.ErrDatabaseUpdateOne(err.Error())
	}

	return nil
}

func (u *usersRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	var model *models.User

	collection := u.db.Collection(model.TableName())

	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errdefs.ErrUserNotFound
		}
		return errdefs.ErrDatabaseDeleteOne(err.Error())
	}

	return nil
}
