package user

import (
	present "ai_writer/internal/presenter/user"
	"ai_writer/internal/router/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	webV1 := router.Group("sien").Group("v1.0").Group("users")
	{
		webV1.POST("", middleware.Verify(), middleware.Transaction(db), control.Create)
		webV1.GET("", middleware.Verify(), control.GetByList)
		webV1.GET(":id", middleware.Verify(), control.GetBySingle)
		webV1.DELETE(":id", middleware.Verify(), middleware.Transaction(db), control.Delete)
		webV1.PATCH(":id", middleware.Verify(), middleware.Transaction(db), control.Update)
		webV1.POST("register", middleware.Transaction(db), control.Create)
	}

	return router
}
