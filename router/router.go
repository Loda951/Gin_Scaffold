package router

import (
	"Gin_Scaffold/logger"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if viper.GetString("mode") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	return router
}
