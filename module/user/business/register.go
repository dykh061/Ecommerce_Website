package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
	"errors"
	"strings"
)

type RegisterRepo interface {
	CreateUser(
		ctx context.Context,
		data *usermodel.UserCreate,
	) error
}

type registerBusiness struct {
	frepo  AciveUEmailFinder
	crepo  RegisterRepo
	hasher PasswordHasher
}

func NewRegisterBusiness(
	frepo AciveUEmailFinder,
	crepo RegisterRepo,
	hasher PasswordHasher,
) *registerBusiness {
	return &registerBusiness{
		frepo:  frepo,
		crepo:  crepo,
		hasher: hasher,
	}
}

func (biz *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {

	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return common.ErrMissingField("name")
	}

	if len(data.Password) < 8 {
		return common.ErrInvalidField("password", "must be at least 8 characters")
	}

	hasUser, err := biz.frepo.FindUserWithEmail(ctx, data.Email)
	if err != nil {
		return common.ErrorDB(err)
	}

	if hasUser != nil {
		if hasUser.IsBanned {
			return common.ErrInvalidState(usermodel.EntityName, "user is banned")
		}
		if hasUser.Status == common.SystemStatusActive {
			return common.ErrEmailAlreadyExists(errors.New("Email already exists"))
		}
		if hasUser.Status == common.SystemStatusDeleted {
			return common.ErrInvalidState(usermodel.EntityName, "deleted")
		}
	}

	hashedPassword, err := biz.hasher.Hash(data.Password)
	if err != nil {
		return err
	}
	data.Password = string(hashedPassword)

	if err := biz.crepo.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
