package srrepository

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

type FindUserWithEmailStorage interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type findUserWithEmailRepo struct {
	storage FindUserWithEmailStorage
}

func NewFindUserWithEmailRepo(storage FindUserWithEmailStorage) *findUserWithEmailRepo {
	return &findUserWithEmailRepo{storage: storage}
}

func (repo *findUserWithEmailRepo) FindUserWithEmail(
	ctx context.Context,
	email string,
) (*usermodel.User, error) {

	return repo.storage.FindDataWithCondition(ctx, map[string]interface{}{
		"email": email,
	})
}
