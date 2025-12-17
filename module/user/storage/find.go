package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,

	// con trỏ ở đây để có thể trả về nil khi không tìm thấy thay vì phải tạo 1 struct rỗng
	// vì khi về struct rỗng sẽ bị mất memory nhiều hơn nil
) (*usermodel.User, error) {

	var data usermodel.User

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
