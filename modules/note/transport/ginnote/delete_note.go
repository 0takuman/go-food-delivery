package ginnote

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	appctx "food-delivery/components/appcontext"
	notebiz "food-delivery/modules/note/biz"
	notestorage "food-delivery/modules/note/storage"
)

func DeleteNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := notestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := notebiz.NewDeleteNoteBiz(store)

		if err := biz.DeleteNote(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
