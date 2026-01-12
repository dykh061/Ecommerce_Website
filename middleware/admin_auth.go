package middleware

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	"OpenMarket/component/tokenprovider/jwt"
	adminstorage "OpenMarket/module/admin/storage"
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	CurrentStaff = "current_staff"
)

// RequireAdminAuth validates admin JWT token and loads staff info
// Similar to RequiredAuthenHeader but for staff/admin accounts
func RequireAdminAuth(appCtx appctx.AppContext) gin.HandlerFunc {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(common.ErrUnauthorized(errors.New("missing or invalid authorization header")))
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(common.ErrUnauthorized(errors.New("invalid or expired token")))
		}

		db := appCtx.GetMainDBConnection()
		storage := adminstorage.NewSQLStore(db)

		staff, err := storage.FindStaffById(c.Request.Context(), payload.UserId)
		if err != nil {
			panic(common.ErrUnauthorized(errors.New("staff not found")))
		}

		if staff.Status == 0 {
			panic(common.ErrForbidden(errors.New("staff account is disabled")))
		}

		// Store staff in context
		c.Set(CurrentStaff, staff)
		c.Next()
	}
}

// RequireRole checks if the authenticated staff has at least one of the required roles
// Must be used AFTER RequireAdminAuth middleware
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		staffRaw, exists := c.Get(CurrentStaff)
		if !exists {
			panic(common.ErrUnauthorized(errors.New("authentication required")))
		}

		staff, ok := staffRaw.(interface {
			HasRole(string) bool
		})
		if !ok {
			panic(common.ErrInternal(errors.New("invalid staff context")))
		}

		// Check if staff has any of the required roles
		hasRequiredRole := false
		for _, role := range roles {
			if staff.HasRole(role) {
				hasRequiredRole = true
				break
			}
		}

		if !hasRequiredRole {
			panic(common.ErrForbidden(errors.New("insufficient permissions")))
		}

		c.Next()
	}
}

// RequireAdmin is a shortcut for RequireRole("admin")
func RequireAdmin() gin.HandlerFunc {
	return RequireRole("admin")
}

// RequireModerator allows admin or moderator roles
func RequireModerator() gin.HandlerFunc {
	return RequireRole("admin", "moderator")
}
