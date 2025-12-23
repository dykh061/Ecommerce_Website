package middleware

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
