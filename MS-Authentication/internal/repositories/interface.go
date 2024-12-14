package repositories

import (
	"context"

	"github.com/charoleizer/thuigsinn/ms-authentication/pkg/proto/authentication"
)

type UsersRepository interface {
	Create(ctx context.Context, users *authentication.RegisterRequest) error
}
