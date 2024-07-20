package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=fadak password=password dbname=postgres host=185.228.236.70 sslmode=disable dbname=fadak port=5432 TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Meeting{}, &Image{}, &Video{}, &Music{})
	if err != nil {
		return
	}

	meetingAPI := NewMeetingAPI(db)

	r := gin.Default()

	r.GET("/api/meetings", meetingAPI.List)
	r.POST("/api/meetings", meetingAPI.Create)

	r.Run() // listen and serve on 0.0.0.0:8080 (default)
}
