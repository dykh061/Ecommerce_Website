package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
)

type FindUserStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type findUserBusiness struct {
	store FindUserStore
}

func NewFindUserBusiness(store FindUserStore) *findUserBusiness {
	return &findUserBusiness{store: store}
}

func (biz *findUserBusiness) FindUser(
	ctx context.Context,
	id int,
) (*usermodel.User, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"id":     id,
		"status": common.SystemStatusActive,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
