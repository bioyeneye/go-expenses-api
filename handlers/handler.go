package handlers

import (
	"github.com/bioyeneye/expenses-api/core/constants"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type ApiRoutes struct {
	Root          *gin.RouterGroup // ''
	ApiRoot       *gin.RouterGroup // 'api/v1'
}

type APIRoutes struct {
	BaseRoutes *ApiRoutes
	DB         *gorm.DB
}

func InitApplication(router *gin.Engine, db *gorm.DB) {
	api := &APIRoutes{
		BaseRoutes: &ApiRoutes{},
		DB:         db,
	}

	api.BaseRoutes.Root = router.Group("/")
	api.BaseRoutes.ApiRoot = router.Group(constants.ApiUrlSuffix)

	router.GET("/_status", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Work!",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "LOST-NOT-FOUND", "message": "I think you are lost, kindly re-route your request rightly MY-GUY."})
	})
}