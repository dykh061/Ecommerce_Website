package userstorage

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
)

// CreateUser là hàm thực thi việc tạo user trong database.
// Hàm này là implement của interface CreateUserStore
// mà business yêu cầu.
func (s *sqlStore) Create(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	return nil
}
