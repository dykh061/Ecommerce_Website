package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

// CreateUser là hàm thực thi việc tạo user trong database.
// Hàm này là implement của interface CreateUserStore
// mà business yêu cầu.
func (s *sqlStore) Create(ctx context.Context, data *usermodel.UserCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
