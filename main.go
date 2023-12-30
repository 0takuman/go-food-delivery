package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	ginnote "food-delivery/modules/note/transport/ginnote"
)

type Notes struct {
	Id     int    `json:"id" gorm:"column:id"`
	Title  string `json:"title" gorm:"column:title"`
	Status bool   `json:"status" gorm:"column:status"`
}

func (Notes) TableName() string {
	return "notes"
}

type NotesUpdate struct {
	Id     int     `json:"id" gorm:"column:id"`
	Title  *string `json:"title" gorm:"column:title"`
	Status bool    `json:"status" gorm:"column:status"`
}

func (NotesUpdate) TableName() string {
	return Notes{}.TableName()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	log.Println(db, err)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/api/v1")

	v1.POST("/notes", ginnote.CreateNote(db))

	v1.GET("/notes/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var note Notes

		if err := db.First(&note, id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": note})
	})

	v1.GET("/notes", func(c *gin.Context) {
		var notes []Notes

		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}

		var pagingData Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}

		if pagingData.Limit <= 0 {
			pagingData.Limit = 10
		}

		if err := db.Offset((pagingData.Page - 1) * pagingData.Limit).Order("id desc").Limit(pagingData.Limit).Find(&notes).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": notes})
	})

	v1.PATCH("/notes/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var note NotesUpdate

		if err := c.ShouldBind(&note); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id=?", id).Updates(note).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": note})
	})

	v1.DELETE("/notes/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(Notes{}.TableName()).Where("id=?", id).Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "success"})
	})

	r.Run()

	// newNote := Notes{Title: "testNote123", Status: true}
	// if err := db.Create(&newNote).Error; err != nil {
	// 	log.Println(err)
	// }

	// var myNote Notes

	// if err := db.First(&myNote).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println(myNote)

	// newTitle := ""
	// updateData := NotesUpdate{Title: &newTitle}
	// if err := db.Where("id=?", 2).Updates(updateData).Error; err != nil {
	// 	log.Println(err)
	// }

	// if err := db.Table(Notes{}.TableName()).Where("id=?", 1).Delete(nil); err != nil {
	// 	log.Println(err)
	// }
}
