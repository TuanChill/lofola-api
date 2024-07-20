package initialize

import (
	"fmt"

	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/routers"
)

func InitRouter() {
	//init router
	r := routers.NewRouter()

	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
