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

func ListNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pagingData.Fullfill()

		var filter notemodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var results []notemodel.Notes
		store := notestorage.NewSQLStore(db)
		biz := notebiz.NewListNoteBiz(store)

		results, err := biz.ListNote(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(results, pagingData, filter))

	}
}
