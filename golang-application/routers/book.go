package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iyashjaat/week2-GL1-CipherSchools.git/database"
	"github.com/iyashjaat/week2-GL1-CipherSchools.git/handler"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := handler.Handler{
		DB: database.GetDB(),
	}
	router.GET("/books", api.GetBooks)
	router.POST("/books", api.SaveBook)
	router.DELETE("/books/:id", api.DeleteBookByID)
	router.PUT("/books/:id", api.UpdateBookByID)
	return router
}
