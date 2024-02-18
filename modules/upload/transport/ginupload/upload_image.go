package ginupload

import (
	"fmt"
	"food-delivery/common"
	appctx "food-delivery/components/appcontext"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename))


		c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
			ID:       0,
			Url:      fmt.Sprintf("http://localhost:9898/static/%s", fileHeader.Filename),
			Width:    0,
			Height:   0,
			CloudName: "local",
			Extension: "png",
		}))
	}
}
