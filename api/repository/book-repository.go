package repository

import (
	"go-playground/top-down-di/api/models"
	"go-playground/top-down-di/internal/db"

	"gorm.io/gorm"
)

// Book 인터페이스 정의
type BookRepository interface {
	GetAllContents() ([]models.Book, error) 
}

// BookRepositoryImpl 함수 작성시 DI 형태로 주입할 DB Conn 정의
type BookRepositoryImpl struct {
	*gorm.DB
}

// 생성자. DB Connection 가지고, BookRepositoryImpl 생성
func NewBookRepositoryImpl() (*BookRepositoryImpl) {
	pg := db.GetDB()
	return &BookRepositoryImpl{pg}
}

// BookRepo 인터페이스 함수 구현
func (pg *BookRepositoryImpl) GetAllContents() ([]models.Book, error) {	
	var books []models.Book
	err := pg.Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}