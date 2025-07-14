package db

import (
	"log"

	"github.com/dayiamin/gin_blog_api/models"
	"gorm.io/driver/sqlite"

	// "github.com/glebarez/sqlite"  //for using pure go for sql
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect(){


	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil{
		log.Fatal("db connections failed")
	}
	db.AutoMigrate(&models.User{},&models.UserProfile{},&models.Post{},&models.PostComment{})
	DB = db

	
}