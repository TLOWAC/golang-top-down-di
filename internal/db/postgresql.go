package db

import (
	"go-playground/top-down-di/api/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() error {

	// DB_URL 환경 변수 값 조회
	dbURL := os.Getenv("DB_URL")
	
	// DB_URL 이 없는 경우 예외처리 및 초기값 할당
	if dbURL == "" {
		dbURL = "postgres://root:root123@localhost:5432/go-local"
	}

	// gorm, postgresql 연동
	pgConn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	
  if err != nil {
		panic("Failed to connect to database")
	}

	pgConn.Migrator().AutoMigrate(&models.Book{})

	db = pgConn
	
	return nil
}

func GetDB() *gorm.DB {
	return db
}