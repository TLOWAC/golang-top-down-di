package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	// Book 엔드포인트 생성
	AddBooksRouter(v1)

	// Auth 엔드포인트 생성
	// AddAuthRoutes(v1)

	return router
}