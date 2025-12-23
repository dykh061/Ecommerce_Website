package userstorage

import (
	usermodel "OpenMarket/module/user/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,

	// con trỏ ở đây để có thể trả về nil khi không tìm thấy thay vì phải tạo 1 struct rỗng
	// vì khi về struct rỗng sẽ bị mất memory nhiều hơn nil
) (*usermodel.User, error) {

	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var data usermodel.User

	if err := db.Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}
