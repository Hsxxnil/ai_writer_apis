package login

import (
	present "ai_writer/internal/presenter/login"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	webV1 := router.Group("sien").Group("v1.0")
	{
		webV1.POST("login", control.Login)
		webV1.POST("refresh", control.Refresh)
	}

	return router
}
