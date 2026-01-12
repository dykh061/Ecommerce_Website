package admingin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"OpenMarket/component/hasher"
	"OpenMarket/component/tokenprovider"
	"OpenMarket/component/tokenprovider/jwt"
	adminmodel "OpenMarket/module/admin/model"
	adminstorage "OpenMarket/module/admin/storage"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type seedCreateStaffRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DevSeedCreateStaff is a DEV-ONLY endpoint to seed a staff account + admin role.
// POST /v1/dev/seed/create-staff
// Body (optional): {"username":"admin","password":"admin123"}
// Guard:
// - Disabled in gin.ReleaseMode
// - Requires env ENABLE_DEV_SEED=true
// - If env DEV_SEED_KEY is set, requires header X-Seed-Key to match
func DevSeedCreateStaff(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.Mode() == gin.ReleaseMode {
			panic(common.ErrForbidden(errors.New("seed endpoint disabled in release mode")))
		}

		if strings.ToLower(strings.TrimSpace(os.Getenv("ENABLE_DEV_SEED"))) != "true" {
			panic(common.ErrForbidden(errors.New("seed endpoint disabled (set ENABLE_DEV_SEED=true)")))
		}

		seedKey := strings.TrimSpace(os.Getenv("DEV_SEED_KEY"))
		if seedKey != "" {
			if strings.TrimSpace(c.GetHeader("X-Seed-Key")) != seedKey {
				panic(common.ErrUnauthorized(errors.New("missing or invalid X-Seed-Key")))
			}
		}

		req := seedCreateStaffRequest{Username: "admin", Password: "admin123"}
		_ = c.ShouldBindJSON(&req) // optional

		req.Username = strings.TrimSpace(req.Username)
		req.Password = strings.TrimSpace(req.Password)
		if req.Username == "" {
			req.Username = "admin"
		}
		if req.Password == "" {
			req.Password = "admin123"
		}

		bcryptHasher := hasher.NewBcryptHasher(12)
		hashedPassword, err := bcryptHasher.Hash(req.Password)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		db := appCtx.GetMainDBConnection()
		store := adminstorage.NewSQLStore(db)
		ctx := c.Request.Context()

		staff, err := store.FindStaffByUsername(ctx, req.Username)
		staffId := 0
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				// Find by username failed unexpectedly
				panic(common.ErrInternal(err))
			}

			// Create new staff
			staffId, err = store.CreateStaff(ctx, &adminmodel.StaffCreate{Username: req.Username}, hashedPassword)
			if err != nil {
				panic(common.ErrInternal(err))
			}
		} else {
			staffId = staff.Id
			// Ensure enabled + password updated to requested
			if err := store.UpdateStaffPasswordAndEnable(ctx, staffId, hashedPassword); err != nil {
				panic(common.ErrInternal(err))
			}
		}

		// Ensure role "admin" exists
		roleId, err := store.EnsureRole(ctx, adminmodel.RoleAdmin)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		// Ensure staff has admin role
		if err := store.EnsureStaffRole(ctx, staffId, roleId); err != nil {
			panic(common.ErrInternal(err))
		}

		// Load full staff (with roles) + return token
		staff, err = store.FindStaffById(ctx, staffId)
		if err != nil {
			panic(common.ErrInternal(err))
		}
		if staff.Status == 0 {
			panic(common.ErrForbidden(errors.New("staff account is disabled")))
		}
		staff.Mask()

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		token, err := tokenProvider.Generate(tokenprovider.TokenPayload{UserId: staff.Id}, 60*60*24*7)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(gin.H{
			"token": token,
			"staff": staff,
		}))
	}
}
