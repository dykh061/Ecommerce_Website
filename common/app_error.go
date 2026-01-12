package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// NewCustom Error đảm bảo dù có root error hay không thì cũng sẽ
// tạo ra một AppError với message và key phù hợp không bao giờ để
// RootErr bị nil
func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)

}

// RootErr sẽ bốc tách e có phải là AppError không, nếu có thì tiếp tục bốc tách
// đến khi nào gặp lỗi gốc (RootErr) không phải AppError nữa thì dừng lại và trả về lỗi gốc đó
// WHY:
//   - Trong Clean Architecture, lỗi thường bị "wrap" nhiều lần
//     (storage -> business -> transport).
//   - Mỗi tầng thêm NGỮ NGHĨA, nhưng lỗi gốc vẫn là DB / system.
//
// Nếu không bóc vỏ:
// - Log sẽ chỉ thấy message nghiệp vụ (ví dụ: "user not found")
// - Dev KHÔNG biết lỗi thật sự là gì (sql.ErrNoRows? timeout? panic?)
//
// RootError đảm bảo:
// - Log luôn ghi nhận lỗi gốc nhất
// - Debug đúng nguyên nhân
// - Không bị nhiễu bởi các wrapper nghiệp vụ
func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

// Apperr implement error của go native interface
// WHY phải gọi RootError()?
//   - Khi log (log.Error(err), fmt.Println(err))
//     ta luôn muốn thấy lỗi gốc thật sự
//   - Tránh log nhầm message nghiệp vụ
//
// Kết quả:
// - AppError vẫn là `error`
// - Log vẫn đúng nguyên nhân
// - Business không cần phân biệt loại error
func (e *AppError) Error() string {
	return e.RootError().Error()
}

// ErrorDB
// DÙNG KHI:
// - Lỗi phát sinh từ tầng database (query fail, connection fail, timeout DB)
// - Không muốn expose chi tiết DB ra client
//
// VÍ DỤ:
// err := db.Create(&user).Error
// return common.ErrorDB(err)
func ErrorDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

// InvalidRequestError
// DÙNG KHI:
// - Request gửi lên sai format
// - Bind JSON / query param bị lỗi
// - Validate body ở tầng transport (Gin)
//
// VÍ DỤ:
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//	    return common.InvalidRequestError(err)
//	}
func InvalidRequestError(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

// ErrInternal
// DÙNG KHI:
// - Lỗi không xác định
// - Panic được recover
// - Fallback error cuối cùng của hệ thống
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong in the server", err.Error(), "ErrInternal")
}

// ErrUserAlreadyHasSeller
// DÙNG KHI:
// - User cố tạo shop mới nhưng đã có shop
//
// HTTP: 400
func ErrUserAlreadyHasSeller(err error) *AppError {
	if err == nil {
		err = errors.New("user already has seller")
	}

	return NewErrorResponse(
		err,
		"user already has a shop",
		err.Error(),
		"ErrUserAlreadyHasSeller",
	)
}

func ErrSellerWasSoftDeleted(err error) *AppError {
	if err == nil {
		err = errors.New("seller was soft deleted")
	}

	return NewErrorResponse(
		err,
		"you are not allowed to create a shop",
		err.Error(),
		"ErrSellerWasSoftDeleted",
	)
}

// ErrCannotListEntity
// DÙNG KHI:
// - Không thể lấy danh sách entity (list / filter / pagination)
// - Lỗi xảy ra ở business hoặc storage
//
// VÍ DỤ:
// users, err := repo.List(ctx)
//
//	if err != nil {
//	    return common.ErrCannotListEntity("User", err)
//	}
func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Canot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

// ErrUnauthorized
// DÙNG KHI:
// - User chưa đăng nhập
// - Token không tồn tại / hết hạn / không hợp lệ
//
// HTTP: 401
func ErrUnauthorized(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusUnauthorized,
		err,
		"unauthorized",
		err.Error(),
		"ErrUnauthorized",
	)
}

// ErrForbidden
// DÙNG KHI:
// - User đã đăng nhập nhưng KHÔNG có quyền
// - Role không đủ (user truy cập admin API)
//
// HTTP: 403
func ErrForbidden(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusForbidden,
		err,
		"forbidden",
		err.Error(),
		"ErrForbidden",
	)
}

// ErrUserNotFound
// DÙNG KHI:
// - Không tìm thấy user theo ID / email
//
// HTTP: 404
func ErrUserNotFound(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusNotFound,
		err,
		"user not found",
		err.Error(),
		"ErrUserNotFound",
	)
}

// ErrEmailAlreadyExists
// DÙNG KHI:
// - Đăng ký user
// - Update email bị trùng
//
// HTTP: 400
func ErrEmailAlreadyExists(err error) *AppError {
	return NewErrorResponse(
		err,
		"email already exists",
		err.Error(),
		"ErrEmailAlreadyExists",
	)
}

// ErrMissingField
// DÙNG KHI:
// - Thiếu field bắt buộc
// - Validate request body / business rule
//
// VÍ DỤ:
//
//	if req.Email == "" {
//	    return common.ErrMissingField("email")
//	}
func ErrMissingField(field string) *AppError {
	return NewErrorResponse(
		errors.New("missing field"),
		fmt.Sprintf("missing field: %s", field),
		field,
		"ErrMissingField",
	)
}

// ErrInvalidField
// DÙNG KHI:
// - Field có giá trị không hợp lệ (email sai format, age < 0)
//
// VÍ DỤ:
//
//	if !isValidEmail(req.Email) {
//	    return common.ErrInvalidField("email", "invalid format")
//	}
func ErrInvalidField(field string, reason string) *AppError {
	return NewErrorResponse(
		errors.New("invalid field"),
		fmt.Sprintf("invalid %s", field),
		reason,
		"ErrInvalidField",
	)
}

// ErrInvalidID
// DÙNG KHI:
// - ID không đúng format (UUID, int <= 0)
//
// VÍ DỤ:
// id, err := strconv.Atoi(c.Param("id"))
//
//	if err != nil {
//	    return common.ErrInvalidID("User")
//	}
func ErrInvalidID(entity string) *AppError {
	return NewErrorResponse(
		errors.New("invalid id"),
		fmt.Sprintf("invalid %s id", strings.ToLower(entity)),
		entity,
		"ErrInvalidID",
	)
}

// ErrPermission
// DÙNG KHI:
// - User KHÔNG được phép thực hiện hành động
// - Bao gồm: banned, role không đủ, scope không hợp lệ
//
// HTTP: 403
func ErrPermission(reason string, err error) *AppError {
	if err == nil {
		err = errors.New("permission denied")
	}

	return NewFullErrorResponse(
		http.StatusForbidden,
		err,
		"permission denied",
		reason,
		"ErrPermission",
	)
}

// ErrEntityNotFound
// DÙNG KHI:
// - Không tìm thấy entity bất kỳ (User, Product, Order)
//
// HTTP: 404
func ErrEntityNotFound(entity string, err error) *AppError {
	return NewFullErrorResponse(
		http.StatusNotFound,
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

// ErrCannotCreateEntity
// DÙNG KHI:
// - Tạo entity thất bại (DB, business rule)
//
// HTTP: 500
func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		fmt.Sprintf("cannot create %s", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

// ErrCannotUpdateEntity
// DÙNG KHI:
// - Update entity thất bại
//
// HTTP: 500
func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		fmt.Sprintf("cannot update %s", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

// ErrCannotReadFile
// DÙNG KHI:
// - Không đọc được file upload
// - io.ReadAll(file) fail
// - file.Open() fail
//
// HTTP: 400 (client gửi file lỗi hoặc không hợp lệ)
func ErrCannotReadFile(err error) *AppError {
	if err == nil {
		err = errors.New("cannot read file")
	}

	return NewErrorResponse(
		err,
		"cannot read file",
		err.Error(),
		"ErrCannotReadFile",
	)
}

// ErrCannotUploadFile
// DÙNG KHI:
// - Upload file lên MinIO / S3 thất bại
// - PutObject error
// - Network lỗi
// - Bucket không tồn tại
// - Permission deny
//
// HTTP: 500
// Lý do:
// - Client gửi file đúng
// - Nhưng hệ thống KHÔNG upload được
func ErrCannotUploadFile(err error) *AppError {
	if err == nil {
		err = errors.New("cannot upload file")
	}

	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"cannot upload file",
		err.Error(),
		"ErrCannotUploadFile",
	)
}

// ErrCannotDeleteEntity
// DÙNG KHI:
// - Delete entity thất bại
//
// HTTP: 500
func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		fmt.Sprintf("cannot delete %s", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

// ErrInvalidState
// DÙNG KHI:
// - Entity ở trạng thái không hợp lệ cho hành động hiện tại
//
// VÍ DỤ:
// - Order đã PAID nhưng vẫn cancel
// - User đã BANNED nhưng vẫn login
func ErrInvalidState(entity string, state string) *AppError {
	return NewErrorResponse(
		errors.New("invalid state"),
		fmt.Sprintf("%s is in invalid state: %s", strings.ToLower(entity), state),
		state,
		"ErrInvalidState",
	)
}

// ErrOperationNotAllowed
// DÙNG KHI:
// - Hành động bị cấm theo business rule
//
// HTTP: 403
func ErrOperationNotAllowed(action string) *AppError {
	return NewFullErrorResponse(
		http.StatusForbidden,
		errors.New("operation not allowed"),
		fmt.Sprintf("%s is not allowed", action),
		action,
		"ErrOperationNotAllowed",
	)
}

// ErrTimeout
// DÙNG KHI:
// - Timeout gọi DB / service bên ngoài
//
// HTTP: 504
func ErrTimeout(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusGatewayTimeout,
		err,
		"request timeout",
		err.Error(),
		"ErrTimeout",
	)
}

// ErrServiceUnavailable
// DÙNG KHI:
// - Gọi service bên ngoài (payment, email, shipping) bị down
//
// HTTP: 503
func ErrServiceUnavailable(service string, err error) *AppError {
	return NewFullErrorResponse(
		http.StatusServiceUnavailable,
		err,
		fmt.Sprintf("%s service unavailable", service),
		err.Error(),
		"ErrServiceUnavailable",
	)
}

// ErrVariantAlreadyExists
// DÙNG KHI:
// - Variant với cùng attribute_value_ids đã tồn tại
//
// HTTP: 400
func ErrVariantAlreadyExists(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusBadRequest,
		err,
		"Biến thể này đã tồn tại",
		err.Error(),
		"ErrVariantAlreadyExists",
	)
}
