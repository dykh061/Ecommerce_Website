package userbusiness

import (
	"context"
	"errors"
)

type ChangePasswordRepo interface {
	UpdatePassword(
		ctx context.Context,
		userID int,
		hashedPassword string,
	) error
}

type changePasswordBiz struct {
	repo       ChangePasswordRepo
	userFinder ActiveUserFinder
	hash       PasswordHasher
}

func NewChangePasswordBiz(
	repo ChangePasswordRepo,
	userFinder ActiveUserFinder,
	hash PasswordHasher,
) *changePasswordBiz {
	return &changePasswordBiz{
		repo:       repo,
		userFinder: userFinder,
		hash:       hash,
	}
}

func (biz *changePasswordBiz) ChangePassword(ctx context.Context, userID int, oldPassword, newPassword string) error {
	user, err := biz.userFinder.FindActiveUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// nghiệp vụ: kiểm tra mật khẩu cũ
	if !biz.hash.Compare(user.Password, oldPassword) {
		return errors.New("old password is incorrect")
	}

	// nghiệp vụ: validate mật khẩu mới
	if len(newPassword) < 8 {
		return errors.New("new password must be at least 8 characters long")
	}

	hashedPassword, err := biz.hash.Hash(newPassword)
	if err != nil {
		return err
	}

	return biz.repo.UpdatePassword(ctx, userID, hashedPassword)
}
