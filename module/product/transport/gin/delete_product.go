package productgin

// func DeleteProduct(appCtx appctx.AppContext) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		db := appCtx.GetMainDBConnection()

// 		id, err := strconv.Atoi(ctx.Param("id"))
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})

// 			return
// 		}
// 		store := productstorage.NewSQLStore(db)
// 		biz := productbusiness.NewDeleteProductBusiness(store)
// 		if err := biz.DeleteProduct(ctx.Request.Context(), id); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
// 	}
// }
