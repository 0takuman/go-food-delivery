package ginuser

import (
	"food-delivery/common"
	appctx "food-delivery/components/appcontext"
	"food-delivery/components/hasher"
	"food-delivery/components/tokenprovider/jwt"
	userbiz "food-delivery/modules/user/biz"
	usermodel "food-delivery/modules/user/model"
	userstorage "food-delivery/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUser usermodel.UserLogin
		if err := c.ShouldBind(&loginUser); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)

		account, err := business.Login(c.Request.Context(), &loginUser)

		if err != nil {
			panic(common.ErrCannotGetEntity(account.Created.GoString(), err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))

	}
}
