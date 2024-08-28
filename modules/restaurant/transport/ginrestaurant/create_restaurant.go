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

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err := c.ShouldBind(&restaurant); err != nil {
			panic(err)
		}

		restaurant.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(appcontext.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &restaurant); err != nil {
			panic(err)
		}

		restaurant.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(restaurant.FakeId.String()))
	}
}
