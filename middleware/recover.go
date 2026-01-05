package middleware

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"fmt"

	"github.com/gin-gonic/gin"
)

// để xứ lí các lỗi panic xảy ra trong quá trình xử lí request mà không được đặt tên rõ ràng
// ví dụ: truy cập vào phần tử của mảng slice bị out of index, truy cập vào trường của con trỏ nil
// những lỗi này sẽ được hệ thống tự động panic và dừng chương trình nếu không được recover
// middleware Recover sẽ giúp recover các lỗi panic này và trả về lỗi internal server error
// thay vì để chương trình bị dừng đột ngột mà không báo lỗi gì cho client
func Recover(c appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {

			// nếu có panic xảy ra thì recover luôn tồn tại giá trị lỗi
			// err chính là giá trị được truyền vào từ hàm panic()
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return
				}
				var root error
				switch e := err.(type) {
				case error:
					root = e
				default:
					root = fmt.Errorf("%v", e)
				}

				appErr := common.ErrInternal(root)
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				return
			}
		}()

		c.Next()
	}
}
