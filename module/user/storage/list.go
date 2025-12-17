package userstorage

// import (
// 	"OpenMarket/common"
// 	usermodel "OpenMarket/module/user/model"
// 	"context"
// )

// func (s *sqlStore) ListDataWithCondition(
// 	context context.Context,
// 	filter *usermodel.Filter,
// 	paging *common.Paging,
// 	moreKeys ...string,
// ) ([]usermodel.User, error) {
// 	var result []usermodel.User

// 	db := s.db.Where("status in (1)")

// 	if f := filter; f != nil {
// 		if f.OwnerId > 0 {
// 			db = db.Where("owner_id = ?", f.OwnerId)
// 		}
// 	}
// 	if err := db.Find(&result).Error; err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
