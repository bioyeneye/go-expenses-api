package utilities

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
)

func SetupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func SetupEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
