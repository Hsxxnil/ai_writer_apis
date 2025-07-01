package article

import (
	present "ai_writer/internal/presenter/article"
	"ai_writer/internal/router/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	v10 := router.Group("ai-writer").Group("v1.0").Group("articles")
	{
		v10.POST("", middleware.Verify(), middleware.Transaction(db), control.Create)
		v10.GET("", middleware.Verify(), control.GetByList)
		v10.GET(":id", middleware.Verify(), control.GetBySingle)
		v10.DELETE(":id", middleware.Verify(), control.Delete)
		v10.PATCH(":id", middleware.Verify(), control.Update)
	}

	return router
}
