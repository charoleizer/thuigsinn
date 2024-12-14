package dtos

import (
	"users/cmd/users"
	"users/internal/errdefs"
	"users/pkg"
)

type ExtendedUpdateRequest struct {
	*users.UpdateRequest
}

func (req *ExtendedUpdateRequest) Validate() error {
	if req.Username != "" {
		err := pkg.IsValidUsername(req.Username)
		if err != nil {
			return errdefs.ErrInvalidUsername(err.Error())
		}
	}

	if req.Email != "" {
		err := pkg.IsValidEmail(req.Email)
		if err != nil {
			return errdefs.ErrInvalidEmail(err.Error())
		}
	}

	return nil
}
