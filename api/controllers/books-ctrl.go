package controllers

import (
	data "go-playground/top-down-di/api/data/response"
	service "go-playground/top-down-di/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BooksController 인터페이스 정의
type BookController interface {
	FindAll(ctx *gin.Context)
}

// BooksController 구현체
type BookControllerImpl struct {
	bookService service.BookService
}

// 생성자
func NewBooksController() BookController {
	bookService := service.NewBookServiceImpl()
	return &BookControllerImpl{bookService}
}

func (bc *BookControllerImpl) FindAll(ctx *gin.Context) {
	books, err := bc.bookService.GetBooks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get books"})
		return
	}

	webResponse := data.Response{
		Code : http.StatusOK,
		Status:"success",
		Data: books,
	}

	ctx.JSON(http.StatusOK, webResponse)
}