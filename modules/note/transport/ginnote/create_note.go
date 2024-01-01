package ginnote

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"food-delivery/common"
	appctx "food-delivery/components/appcontext"
	notebiz "food-delivery/modules/note/biz"
	notemodel "food-delivery/modules/note/model"
	notestorage "food-delivery/modules/note/storage"
)

func CreateNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var note notemodel.Notes

		if err := c.ShouldBind(&note); err != nil {
			panic(err)
		}

		store := notestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := notebiz.NewCreateNoteBiz(store)

		if err := biz.CreateNote(c.Request.Context(), &note); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
