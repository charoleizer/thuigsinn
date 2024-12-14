package dtos

import (
	"users/cmd/users"
	"users/internal/errdefs"
	"users/pkg"
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

	if err := pkg.IsValidUsername(req.Username); err != nil {
		return errdefs.ErrInvalidUsername(err.Error())
	}

	if err := pkg.IsValidEmail(req.Email); err != nil {
		return errdefs.ErrInvalidEmail(err.Error())
	}

	if err := pkg.IsValidPassword(req.Password); err != nil {
		return errdefs.ErrInvalidPassword(err.Error())
	}

	return nil
}

func (req *ExtendedCreateRequest) HashPassword() error {
	hashedPassword, err := pkg.HashPassword(req.Password, PasswordSalt)
	if err != nil {
		return errdefs.ErrHashPassword(err.Error())
	}

	req.Password = hashedPassword
	return nil
}
