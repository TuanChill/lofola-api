package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/controller"
	"github.com/tuanchill/lofola-api/internal/middleware"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

func GroupRouter(r *gin.RouterGroup) {
	group := r.Group("/group")
	{
		group.GET("/info", utils.AsyncHandler(controller.NewGroupController().GetGroup))
		group.GET("/search", utils.AsyncHandler(controller.NewGroupController().SearchGroup))
		private := group.Group("")
		{
			private.Use(middleware.AuthenMiddleware())
			private.POST("/create", utils.AsyncHandler(controller.NewGroupController().CreateGroup))
			private.PUT("/update", utils.AsyncHandler(controller.NewGroupController().UpdateGroup))
		}
	}
}
