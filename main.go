package main

import (
	router "go-playground/top-down-di/api/routes"
	"go-playground/top-down-di/internal/db"
)

func main() {

	// DB Conn
	db.ConnectDB()

	// Router 셋팅
	r := router.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
