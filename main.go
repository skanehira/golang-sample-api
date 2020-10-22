package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/skanehira/golang-sample-api/handler"
	"github.com/skanehira/golang-sample-api/model"
	"github.com/skanehira/golang-sample-api/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&model.UserInfo{}); err != nil {
		log.Fatal(err)
	}

	if err := db.Create(&model.UserInfo{Name: "gorilla"}).Error; err != nil {
		log.Fatal(err)
	}

	us := service.NewServiceUser(db)
	uh := handler.NewUserHandler(us)

	e := echo.New()

	e.GET("/users", uh.Users)

	log.Fatal(e.Start(":80"))
}
