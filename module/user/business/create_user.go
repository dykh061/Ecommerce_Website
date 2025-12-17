package userbusiness

import (
	usermodel "OpenMarket/module/user/model"
	"context"
)

// CreateUserStore là INTERFACE (hợp đồng) mà tầng business yêu cầu.
//
// Business KHÔNG quan tâm storage là MySQL, SQLite hay bất cứ cái gì.
// Business chỉ cần storage đó CÓ KHẢ NĂNG tạo user.
//
// Bất kỳ storage nào muốn được business sử dụng
// thì BẮT BUỘC phải implement interface này.

type CreateUserStore interface {
	Create(context context.Context, data *usermodel.UserCreate) error
}

// createUserBusiness là STRUCT đại diện cho nghiệp vụ "Tạo user".
//
// Struct này KHÔNG làm việc trực tiếp với database.
// Nó chỉ giữ một reference (thông qua interface) tới storage.
//
// store ở đây chính là dependency của business.

type createUserBusiness struct {
	store CreateUserStore
}

// NewCreateUserBusiness là HÀM KHỞI TẠO business.
//
// Storage đã được tạo sẵn ở bên ngoài (main / initializer)
// và được TRUYỀN VÀO cho business (Dependency Injection).
//
// Business không tự tạo storage cho mình.
func NewCreateUserBusiness(store CreateUserStore) *createUserBusiness {
	return &createUserBusiness{store: store}
}

// CreateUser là USE CASE / LOGIC NGHIỆP VỤ tạo user.
//
// Hàm này sẽ được gọi từ handler (Gin / HTTP).
// Business chỉ điều phối nghiệp vụ, không biết DB xử lý thế nào.

func (biz *createUserBusiness) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {

	// Business gọi gián tiếp xuống storage thông qua interface.
	// Không cần biết storage dùng SQL, GORM hay gì khác.
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
