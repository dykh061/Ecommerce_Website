package userbusiness

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

type findUserBusiness struct {
	repo ActiveUserFinder
}

func NewFindUserBusiness(repo ActiveUserFinder) *findUserBusiness {
	return &findUserBusiness{repo: repo}
}

func (biz *findUserBusiness) FindUser(
	ctx context.Context,
	id int,
) (*usermodel.User, error) {
	result, err := biz.repo.FindActiveUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return result, nil
}
