package main

import (
	"fmt"
	"github.com/bioyeneye/expenses-api/core/constants"
	"github.com/bioyeneye/expenses-api/core/middleswares"
	"github.com/bioyeneye/expenses-api/core/utilities"
	"github.com/bioyeneye/expenses-api/db"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"os"
)

//https://github.com/heroku/go-getting-started
func main() {

	utilities.SetupEnvironment()
	utilities.SetupLogOutput()

	dbConfig := db.NewDBConfigFromEnv()
	dbConString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Name, dbConfig.Password)

	print(dbConString)

	server := gin.New()
	server.Use(
		gin.Recovery(),
		middleswares.Logger(),
		middleswares.CORSMiddleware(),
		middleswares.ContentTypeMiddleware(),
		gindump.Dump())

	port := ":" + os.Getenv(constants.Port)
	err := server.Run(port)

	if err != nil {
		panic("Failed to serve project!")
	}
}
