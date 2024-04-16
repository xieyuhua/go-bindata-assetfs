package catchace

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func RouteConfig(r *gin.Engine) {
	//r.StaticFS("/game", http.Dir("static"))
	r.StaticFS("/game", assetFS()) //bindata_assetfs.go
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/game/index.html")
        // 127.0.0.1:80/game/index.html  127.0.0.1:80/game/test/index.html  
	})
  //POST
	r.POST("/CatchAce/create", createCatchAce)
	// GET
	r.GET("/CatchAce", listCatchAce)
}
