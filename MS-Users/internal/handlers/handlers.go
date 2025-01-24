package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/charoleizer/thuigsinn/ms-users/internal/dtos"
	"github.com/charoleizer/thuigsinn/ms-users/internal/repositories"
	"github.com/charoleizer/thuigsinn/ms-users/pkg/brokers"
	"github.com/charoleizer/thuigsinn/ms-users/pkg/proto/authentication"
	"github.com/charoleizer/thuigsinn/ms-users/pkg/proto/users"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CreateTimeout = time.Second * 5
)

type Handlers struct {
	broker nats.JetStreamContext
	db     *mongo.Database
}

func NewHandlers(broker nats.JetStreamContext, db *mongo.Database) *Handlers {
	return &Handlers{
		broker: broker,
		db:     db,
	}
}

func (h *Handlers) Create(ctx context.Context, req *users.CreateRequest) (*users.CreateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, CreateTimeout)
	defer cancel()

	extendedReq := &dtos.ExtendedCreateRequest{CreateRequest: req}
	if err := extendedReq.Validate(); err != nil {
		return nil, err
	}

	if err := extendedReq.HashPassword(); err != nil {
		return nil, err
	}

	usersRepository := repositories.NewUsers(h.db)
	insertedID, err := usersRepository.Create(ctx, *extendedReq)
	if err != nil {
		return nil, err
	}

	registerRequest := &authentication.RegisterRequest{
		Userid:   insertedID.Hex(),
		Email:    extendedReq.Email,
		Username: extendedReq.Username,
		Password: extendedReq.Password,
	}

	registerRequestJSON, err := json.Marshal(registerRequest)
	if err != nil {
		return nil, err
	}

	natsJSConn := brokers.NewNats(h.broker)
	err = natsJSConn.Publish("users.created", registerRequestJSON)
	if err != nil {
		return nil, err
	}

	// go func(id primitive.ObjectID, req *dtos.ExtendedCreateRequest) {
	// 	authenticationClient, err := services.NewAuthenticationClient(ctx, "localhost:8081")
	// 	if err != nil {
	// 		usersRepository.SetStatus(context.Background(), id, users.Status_Failed)
	// 		return
	// 	}
	// 	defer authenticationClient.AuthenticationClose()

	// 	err = authenticationClient.AuthenticationRegister(context.Background(), &authentication.RegisterRequest{
	// 		Userid:   id.Hex(),
	// 		Email:    req.Email,
	// 		Username: req.Username,
	// 		Password: req.Password,
	// 	})
	// 	if err != nil {
	// 		usersRepository.SetStatus(context.Background(), id, users.Status_Failed)
	// 		return
	// 	}

	// 	usersRepository.SetStatus(context.Background(), id, users.Status_Completed)

	// }(insertedID, extendedReq)

	return &users.CreateResponse{
		Id: insertedID.Hex(),
	}, nil
}

func (h *Handlers) Read(ctx context.Context, req *users.ReadRequest) (*users.ReadResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, CreateTimeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	usersRepository := repositories.NewUsers(h.db)
	user, err := usersRepository.Read(ctx, objectID)
	if err != nil {
		return nil, err
	}

	return &users.ReadResponse{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (h *Handlers) Update(ctx context.Context, req *users.UpdateRequest) (*users.UpdateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, CreateTimeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	extendedReq := &dtos.ExtendedUpdateRequest{UpdateRequest: req}
	if err := extendedReq.Validate(); err != nil {
		return nil, err
	}

	usersRepository := repositories.NewUsers(h.db)
	err = usersRepository.Update(ctx, objectID, *extendedReq)
	if err != nil {
		return nil, err
	}

	return &users.UpdateResponse{
		Id: req.Id,
	}, nil
}

func (h *Handlers) Delete(ctx context.Context, req *users.DeleteRequest) (*users.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, CreateTimeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	usersRepository := repositories.NewUsers(h.db)
	err = usersRepository.Delete(ctx, objectID)
	if err != nil {
		return nil, err
	}

	return &users.DeleteResponse{
		Id: req.Id,
	}, nil
}
