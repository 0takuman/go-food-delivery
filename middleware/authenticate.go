package middleware

import (
	"errors"
	"fmt"
	"food-delivery/common"
	appctx "food-delivery/components/appcontext"
	"food-delivery/components/tokenprovider/jwt"
	userstorage "food-delivery/modules/user/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authentication header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeader(s string) (string, error) {
	parts := strings.Split(s, " ")

	return parts[1], nil
}

func Requireauth(appctx appctx.AppContext) func(c *gin.Context) {
	tokenprovider := jwt.NewTokenJWTProvider(appctx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appctx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)

		payload, err := tokenprovider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
