package ginuser

import (
	"food-delivery/common"
	appctx "food-delivery/components/appcontext"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
