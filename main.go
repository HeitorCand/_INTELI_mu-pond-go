package main

import (
	"mu-pond-go/domain"
	"mu-pond-go/handler"
	"mu-pond-go/repository"
	"mu-pond-go/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("figurinhas.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	db.AutoMigrate(&domain.Figurinha{})

	repo := repository.NewFigurinhaRepository(db)
	svc := service.NewFigurinhaService(repo)
	h := handler.NewFigurinhaHandler(svc)

	r := gin.Default()

	r.POST("/figurinha", h.Create)
	r.GET("/figurinha", h.List)
	r.GET("/figurinha/:id", h.GetByID)
	r.PUT("/figurinha/:id", h.Update)
	r.DELETE("/figurinha/:id", h.Delete)

	r.Run(":8080")
}
