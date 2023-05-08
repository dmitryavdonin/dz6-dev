package delivery

import (
	authRequest "auth/internal/delivery/http/auth"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (d *Delivery) SignIn(c *gin.Context) {
	request := authRequest.SignInRequest{}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionId, err := d.services.Auth.SignIn(context.Background(), request.Login, request.Pass)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		d.logger.Debug("http/auth: SignIn(): FAILED! login = " + request.Login + "; pass = " + request.Pass + "; err = " + err.Error())
		return
	}

	c.SetCookie("session_id", sessionId.String(), int(time.Hour), "/", "", false, true)

	c.JSON(http.StatusOK, authRequest.SignInResponse{SessionId: sessionId})
}

func (d *Delivery) SignOut(c *gin.Context) {
	sessionIdString, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionId, err := uuid.Parse(sessionIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = d.services.Auth.DeleteSessionById(context.Background(), sessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//deletes cookie
	c.SetCookie("session_id", sessionId.String(), -1, "/", "", false, true)

	c.Status(http.StatusOK)
}
