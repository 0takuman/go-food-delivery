package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	appctx "food-delivery/components/appcontext"
	ginnote "food-delivery/modules/note/transport/ginnote"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_CONNECTION")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")

	v1.POST("/notes", ginnote.CreateNote(appCtx))

	v1.GET("/notes", ginnote.ListNote(appCtx))

	// v1.GET("/notes/:id", func(c *gin.Context) {
	// 	id, err := strconv.Atoi(c.Param("id"))
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	var note Notes

	// 	if err := db.First(&note, id).Error; err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"data": note})
	// })

	// v1.GET("/notes", func(c *gin.Context) {
	// 	var notes []Notes

	// 	type Paging struct {
	// 		Page  int `json:"page" form:"page"`
	// 		Limit int `json:"limit" form:"limit"`
	// 	}

	// 	var pagingData Paging

	// 	if err := c.ShouldBind(&pagingData); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	if pagingData.Page <= 0 {
	// 		pagingData.Page = 1
	// 	}

	// 	if pagingData.Limit <= 0 {
	// 		pagingData.Limit = 10
	// 	}

	// 	if err := db.Offset((pagingData.Page - 1) * pagingData.Limit).Order("id desc").Limit(pagingData.Limit).Find(&notes).Error; err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"data": notes})
	// })

	// v1.PATCH("/notes/:id", func(c *gin.Context) {
	// 	id, err := strconv.Atoi(c.Param("id"))
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	var note NotesUpdate

	// 	if err := c.ShouldBind(&note); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	if err := db.Where("id=?", id).Updates(note).Error; err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"data": note})
	// })

	v1.DELETE("/notes/:id", ginnote.DeleteNote(appCtx))

	r.Run()

}
