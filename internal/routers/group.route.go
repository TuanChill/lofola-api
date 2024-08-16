package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/middleware"
	"github.com/tuanchill/lofola-api/internal/wire"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

func GroupRouter(r *gin.RouterGroup) {
	groupController, _ := wire.InitGroupRouterHandler()

	group := r.Group("/group")
	{
		group.GET("/info", utils.AsyncHandler(groupController.GetGroup))
		group.GET("/search", utils.AsyncHandler(groupController.SearchGroup))
		private := group.Group("")
		{
			private.Use(middleware.AuthenMiddleware())
			private.POST("/join", utils.AsyncHandler(groupController.JoinGroup))
			private.POST("/leave", utils.AsyncHandler(groupController.LeaveGroup))
			private.POST("/create", utils.AsyncHandler(groupController.CreateGroup))
			private.PUT("/update", utils.AsyncHandler(groupController.UpdateGroup))
		}
	}
}
