package db

import (
	"github.com/bioyeneye/expenses-api/core"
	"github.com/bioyeneye/expenses-api/core/constants"
	"github.com/bioyeneye/expenses-api/core/utilities"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func NewDBConfig(host string, username string, password string, name string, port string) *core.DBConfig {
	return &core.DBConfig{Host: host, Username: username, Password: password, Name: name, Port: port}
}

func NewDBConfigFromEnv() *core.DBConfig {
	return &core.DBConfig{
		Host:     utilities.GetEnv(constants.DatabaseServer, ""),
		Username: utilities.GetEnv(constants.DatabaseUsernameKey, ""),
		Password: utilities.GetEnv(constants.DatabasePasswordKey, ""),
		Name:     utilities.GetEnv(constants.DatabaseNameKey, ""),
		Port:     utilities.GetEnv(constants.DatabasePortKey, ""),
	}
}

func SetupDbModels(dialect string, dbConString string, entities []interface{}) (*gorm.DB, error) {

	db, err := gorm.Open(dialect, dbConString)
	if err != nil {
		return nil, err
	}

	if len(entities) > 0 {
		db.AutoMigrate(entities...)
	}
	return db, err
}
