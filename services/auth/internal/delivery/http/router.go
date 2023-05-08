package delivery

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (d *Delivery) initRouter() *gin.Engine {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true

	var router = gin.Default()

	router.Use(cors.New(corsConfig))
	router.LoadHTMLFiles("./internal/templates/sign-in.html", "./internal/templates/sign-up.html")
	router.Static("/.well-known", ".well-known")

	router.POST("/sign-in", d.SignIn)
	router.GET("/sign-out", d.SignOut)
	router.GET("/session/:id", d.ReadSessionById)
	router.GET("/session/cookie/*url", d.ReadSessionByCookie)

	return router
}
