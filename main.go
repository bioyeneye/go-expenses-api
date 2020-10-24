package main

import (
	"fmt"
	"github.com/bioyeneye/expenses-api/core/constants"
	"github.com/bioyeneye/expenses-api/core/middleswares"
	"github.com/bioyeneye/expenses-api/core/utilities"
	"github.com/bioyeneye/expenses-api/db"
	"github.com/bioyeneye/expenses-api/db/entities"
	"github.com/bioyeneye/expenses-api/handlers"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"os"
)

//https://github.com/heroku/go-getting-started
//https://medium.com/@sathishvj/web-handlers-and-middleware-in-golang-2706c2ecfb75
func main() {

	utilities.SetupEnvironment()
	utilities.SetupLogOutput()

	dbConString := ""
	env := os.Getenv(constants.Environment)
	if env == "dev" {
		dbConfig := db.NewDBConfigFromEnv()
		dbConString = fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
			dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Name, dbConfig.Password)
	}else{
		dbConString = os.Getenv(constants.DatabaseUrl)
	}

	dbEntities := []interface{} {
		&entities.Users{},
	}

	dbInstance, dbErr := db.SetupDbModels("postgres", dbConString, dbEntities)
	if dbErr != nil {
		panic(dbErr.Error())
	}

	dbInstance.LogMode(true)
	defer dbInstance.Close()

	server := gin.New()
	server.Use(
		gin.Recovery(),
		middleswares.Logger(),
		middleswares.CORSMiddleware(),
		middleswares.ContentTypeMiddleware(),
		gindump.Dump())

	handlers.InitApplication(server, dbInstance)

	port := ":" + os.Getenv(constants.Port)
	err := server.Run(port)

	if err != nil {
		panic("Failed to serve project!")
	}
}
