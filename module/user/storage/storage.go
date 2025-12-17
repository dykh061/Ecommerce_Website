package userstorage

import "gorm.io/gorm"

// sqlStore là struct cài đặt các phương thức lưu trữ dữ liệu sử dụng GORM.
type sqlStore struct {
	db *gorm.DB
}

// NewSQLStore là hàm khởi tạo một sqlStore với kết nối database đã cho.
func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
