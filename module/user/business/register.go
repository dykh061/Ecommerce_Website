package userbusiness

import (
	"OpenMarket/common"
	usermodel "OpenMarket/module/user/model"
	"context"
	"errors"
	"strings"
)

// CreateUserStore là INTERFACE (hợp đồng) mà tầng business yêu cầu.
//
// Business KHÔNG quan tâm storage là MySQL, SQLite hay bất cứ cái gì.
// Business chỉ cần storage đó CÓ KHẢ NĂNG tạo user.
//
// Bất kỳ storage nào muốn được business sử dụng
// thì BẮT BUỘC phải implement interface này.

type RegisterStore interface {
	Create(ctx context.Context, data *usermodel.UserCreate) error
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

// createUserBusiness là STRUCT đại diện cho nghiệp vụ "Tạo user".
//
// Struct này KHÔNG làm việc trực tiếp với database.
// Nó chỉ giữ một reference (thông qua interface) tới storage.
//
// store ở đây chính là dependency của business.

type registerBusiness struct {
	store  RegisterStore
	hasher PasswordHasher
}

// NewCreateUserBusiness là HÀM KHỞI TẠO business.
//
// Storage đã được tạo sẵn ở bên ngoài (main / initializer)
// và được TRUYỀN VÀO cho business (Dependency Injection).
//
// Business không tự tạo storage cho mình.
func NewRegisterBusiness(store RegisterStore, hasher PasswordHasher) *registerBusiness {
	return &registerBusiness{store: store, hasher: hasher}
}

// CreateUser là USE CASE / LOGIC NGHIỆP VỤ tạo user.
//
// Hàm này sẽ được gọi từ handler (Gin / HTTP).
// Business chỉ điều phối nghiệp vụ, không biết DB xử lý thế nào.

func (biz *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {

	// Business gọi gián tiếp xuống storage thông qua interface.
	// Không cần biết storage dùng SQL, GORM hay gì khác.
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return common.ErrMissingField("name")
	}

	if len(data.Password) < 8 {
		return common.ErrInvalidField("password", "must be at least 8 characters")
	}

	hasUser, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"email": data.Email,
	})
	if err != nil {
		return common.ErrorDB(err)
	}

	if hasUser != nil {
		if hasUser.IsBanned {
			return common.ErrInvalidState(usermodel.EntityName, "user is banned")
		}
		if hasUser.Status == common.SystemStatusActive {
			return common.ErrEmailAlreadyExists(errors.New("Email already exists"))
		}
		if hasUser.Status == common.SystemStatusDeleted {
			return common.ErrInvalidState(usermodel.EntityName, "deleted")
		}
	}

	hashedPassword, err := biz.hasher.Hash(data.Password)
	if err != nil {
		return err
	}
	data.Password = string(hashedPassword)

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
