package srrepository

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
)

type FindUserStorage interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type findUserRepo struct {
	storage FindUserStorage
}

func NewFindUserRepo(storage FindUserStorage) *findUserRepo {
	return &findUserRepo{storage: storage}
}

func (repo *findUserRepo) FindActiveUserByID(
	ctx context.Context,
	userID int,
) (*usermodel.User, error) {

	return repo.storage.FindDataWithCondition(ctx, map[string]interface{}{
		"id":     userID,
		"status": common.SystemStatusActive,
	})
}
