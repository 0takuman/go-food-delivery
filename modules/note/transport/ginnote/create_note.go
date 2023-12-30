package ginnote

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	notebiz "food-delivery/modules/note/biz"
	notemodel "food-delivery/modules/note/model"
	notestorage "food-delivery/modules/note/storage"
)

func CreateNote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var note notemodel.Notes

		if err := c.ShouldBind(&note); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := notestorage.NewSQLStore(db)
		biz := notebiz.NewCreateNoteBiz(store)

		if err := biz.CreateNote(c.Request.Context(), &note); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": note})
	}
}
