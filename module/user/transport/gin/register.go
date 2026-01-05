package ginuser

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"OpenMarket/component/hasher"
	validateemail "OpenMarket/component/validate"
	userbusiness "OpenMarket/module/user/business"
	usermodel "OpenMarket/module/user/model"
	srrepository "OpenMarket/module/user/repository"
	userstorage "OpenMarket/module/user/storage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser là hàm xử lý HTTP request tạo user.
//
// Hàm này sẽ:
// 1. Nhận request từ client (Gin)
// 2. Gọi business để thực hiện nghiệp vụ tạo user
// 3. Trả về response cho client (Gin)
//
// Hàm này KHÔNG làm việc trực tiếp với database.
// Nó chỉ điều phối luồng dữ liệu giữa client và business.
func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		if !validateemail.ValidateEmail(data.Email) {
			panic(errors.New("email is not valid"))
		}
		store := userstorage.NewSQLStore(db)
		hasher := hasher.NewBcryptHasher(bcrypt.DefaultCost)
		frepo := srrepository.NewFindUserWithEmailRepo(store)
		crepo := srrepository.NewCreateUserRepo(store)
		biz := userbusiness.NewRegisterBusiness(frepo, crepo, hasher)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
