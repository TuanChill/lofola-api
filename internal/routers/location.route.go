package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/wire"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

func LocationRouter(r *gin.RouterGroup) {
	locationController, _ := wire.InitLocationRouterHandler()

	location := r.Group("/location")
	{
		location.GET("/district", utils.AsyncHandler(locationController.GetDistrict))
		location.GET("/province", utils.AsyncHandler(locationController.GetProvince))
		location.GET("/ward", utils.AsyncHandler(locationController.GetWard))
	}
}
