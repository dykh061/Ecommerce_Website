package adminbusiness

import (
	"OpenMarket/common"
	"OpenMarket/component/hasher"
	"OpenMarket/component/tokenprovider"
	adminmodel "OpenMarket/module/admin/model"
	"context"
	"errors"
)

type LoginStaffStore interface {
	FindStaffByUsername(ctx context.Context, username string) (*adminmodel.Staff, error)
}

type loginStaffBusiness struct {
	store         LoginStaffStore
	tokenProvider tokenprovider.Provider
	hasher        *hasher.BcryptHasher
	expiry        int
}

func NewLoginStaffBusiness(
	store LoginStaffStore,
	tokenProvider tokenprovider.Provider,
	hasher *hasher.BcryptHasher,
	expiry int,
) *loginStaffBusiness {
	return &loginStaffBusiness{
		store:         store,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginStaffBusiness) Login(
	ctx context.Context,
	data *adminmodel.StaffLogin,
) (*tokenprovider.Token, *adminmodel.Staff, error) {
	staff, err := biz.store.FindStaffByUsername(ctx, data.Username)
	if err != nil {
		return nil, nil, common.ErrUnauthorized(errors.New("invalid credentials"))
	}

	if staff.Status == 0 {
		return nil, nil, common.ErrForbidden(errors.New("account is disabled"))
	}

	// Verify password
	if !biz.hasher.Compare(staff.Password, data.Password) {
		return nil, nil, common.ErrUnauthorized(errors.New("invalid credentials"))
	}

	// Generate token with staff payload
	payload := tokenprovider.TokenPayload{
		UserId: staff.Id,
	}

	token, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, nil, common.ErrInternal(err)
	}

	staff.Mask()
	return token, staff, nil
}
