package ginuser

import (
	"food-delivery/common"
	appctx "food-delivery/components/appcontext"
	"food-delivery/components/hasher"
	userbiz "food-delivery/modules/user/biz"
	usermodel "food-delivery/modules/user/model"
	userstorage "food-delivery/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(usermodel.EntityName, err))
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
