package http

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

func (del *Delivery) initRouter() *gin.Engine {
	if viper.GetBool("IS_PRODUCTION") {
		switch strings.ToUpper(strings.TrimSpace(viper.GetString("LOG_LEVEL"))) {
		case "DEBUG":
			gin.SetMode(gin.DebugMode)
		default:
			gin.SetMode(gin.ReleaseMode)

		}
	} else {
		gin.SetMode(gin.DebugMode)
	}
	var router = gin.New()
	router.Use(checkAuthentication)

	router.GET("/contacts", del.handleGetContacts)
	router.POST("/contacts", del.handleCreateContact)
	router.PUT("/contacts/:id", del.handleUpdateContact)
	router.DELETE("/contacts/:id", del.handleDeleteContact)

	router.GET("/groups", del.handleGetGroups)
	router.POST("/groups", del.handleCreateGroup)
	router.PUT("/groups/:id", del.handleUpdateGroup)
	router.DELETE("/groups/:id", del.handleDeleteGroup)

	return router
}
