package dtos

import (
	"github.com/charoleizer/thuigsinn/ms-users/internal/errdefs"
	"github.com/charoleizer/thuigsinn/ms-users/pkg/proto/users"
	"github.com/charoleizer/thuigsinn/ms-users/pkg/utils"
)

type ExtendedUpdateRequest struct {
	*users.UpdateRequest
}

func (req *ExtendedUpdateRequest) Validate() error {
	if req.Username != "" {
		err := utils.IsValidUsername(req.Username)
		if err != nil {
			return errdefs.ErrInvalidUsername(err.Error())
		}
	}

	if req.Email != "" {
		err := utils.IsValidEmail(req.Email)
		if err != nil {
			return errdefs.ErrInvalidEmail(err.Error())
		}
	}

	return nil
}
