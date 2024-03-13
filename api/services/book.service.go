package services

import (
	"go-playground/top-down-di/api/models"
	"go-playground/top-down-di/api/repository"
)

// BookService 에 정의할 함수 인터페이스 정의
type BookService interface {
	GetBooks() ([]models.Book, error) 
	// GetBook(c *gin.Context)
	// AddBook(c *gin.Context)
	// UpdateBook(c *gin.Context)
	// DeleteBook(c *gin.Context)
}

// BookService 구현체
type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

// 생성자
func NewBookServiceImpl() (BookService) {
	bookRepository := repository.NewBookRepositoryImpl()
	return &BookServiceImpl{bookRepository}
}

func (bs *BookServiceImpl) GetBooks() ([]models.Book, error) {
	contents, err := bs.bookRepository.GetAllContents()

	if err != nil {
		return nil, err
	}

	return contents, nil
}