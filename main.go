package main

import (
	"food-delivery/middleware"
	"food-delivery/modules/restaurant/transport/ginrestaurant"
	"food-delivery/modules/upload/transport/ginupload"
	"food-delivery/modules/user/transport/ginuser"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	appctx "food-delivery/components/appcontext"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_CONNECTION")
	secretKey := os.Getenv("SECRET_KEY")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	appCtx := appctx.NewAppContext(db, secretKey)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.Static("/static", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	v1.POST("/upload", ginupload.UploadImage(appCtx))
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.Profile(appCtx))
	restaurant := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurant.GET("/:id", ginrestaurant.ListRestaurant(appCtx))

	// v1.POST("/notes", ginnote.CreateNote(appCtx))

	// v1.GET("/notes", ginnote.ListNote(appCtx))

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

	// v1.DELETE("/notes/:id", ginnote.DeleteNote(appCtx))

	r.Run(":9898")
}
