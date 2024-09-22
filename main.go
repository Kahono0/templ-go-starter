package main

import (
	"os"
	"starter/db"
	"starter/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func dbUrl() string {
	return os.Getenv("DATABASE_URL")
}

func main() {
	db.Connect(dbUrl())

	r := gin.Default()

	router.SetUpRoutes(r)
	r.Run()
}

