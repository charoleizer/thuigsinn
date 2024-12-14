package dtos

import (
	"github.com/charoleizer/thuigsinn/ms-users/internal/errdefs"
	"github.com/charoleizer/thuigsinn/ms-users/pkg/proto/users"
	"github.com/charoleizer/thuigsinn/ms-users/pkg/utils"
)

type ExtendedCreateRequest struct {
	*users.CreateRequest
}

const (
	PasswordSalt = "user_password_salt"
)

func (req *ExtendedCreateRequest) Validate() error {
	if req.Username == "" {
		return errdefs.ErrUsernameRequired
	}

	if req.Email == "" {
		return errdefs.ErrEmailRequired
	}

	if req.Password == "" {
		return errdefs.ErrPasswordRequired
	}

	if err := utils.IsValidUsername(req.Username); err != nil {
		return errdefs.ErrInvalidUsername(err.Error())
	}

	if err := utils.IsValidEmail(req.Email); err != nil {
		return errdefs.ErrInvalidEmail(err.Error())
	}

	if err := utils.IsValidPassword(req.Password); err != nil {
		return errdefs.ErrInvalidPassword(err.Error())
	}

	return nil
}

func (req *ExtendedCreateRequest) HashPassword() error {
	hashedPassword, err := utils.HashPassword(req.Password, PasswordSalt)
	if err != nil {
		return errdefs.ErrHashPassword(err.Error())
	}

	req.Password = hashedPassword
	return nil
}
