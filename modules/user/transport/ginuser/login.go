package ginuser

import (
	"errors"
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
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 30*24*60*60) // 30 days in seconds

		account, err := business.Login(c.Request.Context(), &loginUser)
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.ErrCannotGetEntity("account", err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}


func Logout(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if len(authorizationHeader) > 7 && authorizationHeader[0:7] == "Bearer " {
			token := authorizationHeader[7:]
			tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
			if _, err := tokenProvider.Validate(token); err != nil {
				c.JSON(http.StatusUnauthorized, common.ErrInvalidRequest(err))
				return
			}

			c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]string{
				"message": "logged out successfully",
			}))
		} else {
			c.JSON(http.StatusUnauthorized, common.ErrInvalidRequest(errors.New("missing or invalid JWT token")))
		}
	}
}
