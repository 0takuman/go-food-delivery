package ginrestaurant

import (
	"food-delivery/common"
	"net/http"

	appCtx "food-delivery/components/appcontext"
	restaurantbiz "food-delivery/modules/restaurant/biz"
	restaurantmodel "food-delivery/modules/restaurant/model"
	restaurantstorage "food-delivery/modules/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appcontext appCtx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var restaurant restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&restaurant); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSQLStore(appcontext.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &restaurant); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
