package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/configs/common/constants"
)

func StaticFilesRouter(r *gin.Engine) {
	r.Static(constants.UploadDir, "."+constants.UploadDir)
}
