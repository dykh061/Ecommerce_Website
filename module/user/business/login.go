package userbusiness

import (
	"OpenMarket/common"
	"OpenMarket/component/tokenprovider"
	usermodel "OpenMarket/module/user/model"
	"context"
	"errors"
)

type loginBusiness struct {
	repoUser      AciveUEmailFinder
	tokenProvider tokenprovider.Provider
	hasher        PasswordHasher
	expiry        int
}

func NewLoginBusiness(
	repoUser AciveUEmailFinder,
	tokenProvider tokenprovider.Provider,
	hasher PasswordHasher,
	expiry int,
) *loginBusiness {
	return &loginBusiness{
		repoUser:      repoUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}
func (biz *loginBusiness) Login(
	ctx context.Context,
	data *usermodel.UserLogin,
) (*tokenprovider.Token, error) {
	user, err := biz.repoUser.FindUserWithEmail(ctx, data.Email)

	if err != nil || user == nil {
		return nil, errors.New("Email Or Password Invalid")
	}

	if user.Status == common.SystemStatusDeleted {
		return nil, common.ErrPermission("Account is disabled", errors.New("Account is disabled"))
	}

	if user.IsBanned {
		return nil, common.ErrPermission("Account is disabled", errors.New("Account is disabled"))
	}

	if !biz.hasher.Compare(user.Password, data.Password) {
		return nil, errors.New("Email Or Password Invalid")
	}
	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return accessToken, nil
}
