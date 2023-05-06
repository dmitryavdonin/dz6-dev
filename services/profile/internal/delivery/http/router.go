package delivery

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (d *Delivery) initRouter() *gin.Engine {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true

	var router = gin.Default()

	router.Use(cors.New(corsConfig))

	//var router = gin.New()
	router.POST("/sign-up", d.CreateUser)
	router.POST("/user/creds", d.ReadUserByCredetinals)
	router.Use(d.checkAuth)
	d.routerUser(router.Group("/user"))

	return router
}

func (d *Delivery) routerUser(router *gin.RouterGroup) {
	router.GET("/:id", d.ReadUserById)
	router.PUT("/:id", d.UpdateUser)
	router.DELETE("/:id", d.DeleteUserById)
}

func (d *Delivery) checkAuth(c *gin.Context) {
	token := c.GetHeader("x-auth-token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "token not provided")
		return
	}

	result, err := d.tm.Parse(token)
	if err != nil {
		d.logger.Error("error parsing token - %s: %s", token, err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, "token parse error")
		return
	}

	setAuthInfoToContext(c, result.UserID, result.Login)

	c.Next()
}

func setAuthInfoToContext(c *gin.Context, userId string, login string) {
	c.Set("userId", userId)
	c.Set("userName", login)
}
