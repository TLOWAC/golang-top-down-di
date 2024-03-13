package router

import (
	"go-playground/top-down-di/api/controllers"

	"github.com/gin-gonic/gin"
)

func AddBooksRouter(rg *gin.RouterGroup) {
	bookRouter := rg.Group("/books")
	bookCtrl := controllers.NewBooksController()

	bookRouter.GET("", bookCtrl.FindAll)

	// content.GET("/:id", h.GetContent)
	// content.POST("", h.AddContent)
	// content.PATCH("/:id", h.UpdateContent)
	// content.DELETE("/:id", h.DeleteContent)
}