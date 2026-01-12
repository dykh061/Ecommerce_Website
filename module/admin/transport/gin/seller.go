package admingin

import (
	"OpenMarket/common"
	"OpenMarket/component/appctx"
	adminbusiness "OpenMarket/module/admin/business"
	adminmodel "OpenMarket/module/admin/model"
	adminstorage "OpenMarket/module/admin/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListSellers returns paginated list of sellers for admin panel
// GET /v1/admin/sellers
func ListSellers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter adminmodel.SellerFilter
		var paging common.Paging

		if err := c.ShouldBindQuery(&filter); err != nil {
			panic(common.InvalidRequestError(err))
		}

		if err := c.ShouldBindQuery(&paging); err != nil {
			panic(common.InvalidRequestError(err))
		}
		paging.Fulfill()

		db := appCtx.GetMainDBConnection()
		store := adminstorage.NewSQLStore(db)
		biz := adminbusiness.NewListSellersBusiness(store)

		sellers, err := biz.ListSellers(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(sellers, paging, filter))
	}
}

// GetSeller returns a seller detail for admin panel
// GET /v1/admin/sellers/:id
func GetSeller(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}

		db := appCtx.GetMainDBConnection()
		store := adminstorage.NewSQLStore(db)
		biz := adminbusiness.NewGetSellerBusiness(store)

		seller, err := biz.GetSeller(c.Request.Context(), int(uid.GetLoacalID()))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(seller))
	}
}

// UpdateSellerStatus updates seller status (lock/unlock)
// PATCH /v1/admin/sellers/:id/status
func UpdateSellerStatus(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}

		var data adminmodel.SellerStatusUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.InvalidRequestError(err))
		}

		db := appCtx.GetMainDBConnection()
		store := adminstorage.NewSQLStore(db)
		biz := adminbusiness.NewUpdateSellerStatusBusiness(store)

		if err := biz.UpdateStatus(c.Request.Context(), int(uid.GetLoacalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

// LockSeller locks a seller (shortcut for status=0)
// POST /v1/admin/sellers/:id/lock
func LockSeller(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}

		var body struct {
			Reason string `json:"reason"`
		}
		_ = c.ShouldBindJSON(&body)

		db := appCtx.GetMainDBConnection()
		store := adminstorage.NewSQLStore(db)
		biz := adminbusiness.NewUpdateSellerStatusBusiness(store)

		data := &adminmodel.SellerStatusUpdate{
			Status: adminmodel.SellerStatusInactive,
			Reason: body.Reason,
		}

		if err := biz.UpdateStatus(c.Request.Context(), int(uid.GetLoacalID()), data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

// UnlockSeller unlocks a seller (shortcut for status=1)
// POST /v1/admin/sellers/:id/unlock
func UnlockSeller(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.InvalidRequestError(err))
		}

		db := appCtx.GetMainDBConnection()
		store := adminstorage.NewSQLStore(db)
		biz := adminbusiness.NewUpdateSellerStatusBusiness(store)

		data := &adminmodel.SellerStatusUpdate{
			Status: adminmodel.SellerStatusActive,
		}

		if err := biz.UpdateStatus(c.Request.Context(), int(uid.GetLoacalID()), data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
