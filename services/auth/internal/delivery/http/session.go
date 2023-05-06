package delivery

import (
	"context"
	"net/http"

	jsonSession "auth/internal/delivery/http/session"
	"auth/internal/domain/session"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (d *Delivery) ReadSessionById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := d.services.Auth.ReadSessionById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d.toResponseSession(session))
}

func (d *Delivery) ReadSessionByCookie(c *gin.Context) {
	sessionCookieString, err := c.Request.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(sessionCookieString.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	session, err := d.services.Auth.ReadSessionById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Header("x-username", session.Login())
	c.Header("x-auth-token", session.Token())

	c.JSON(http.StatusOK, d.toResponseSession(session))
}

func (d *Delivery) DeleteSessionById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = d.services.Auth.DeleteSessionById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (d *Delivery) toResponseSession(session *session.Session) *jsonSession.SessionResponse {
	return &jsonSession.SessionResponse{
		Id:         session.Id(),
		Login:      session.Login(),
		Token:      session.Token(),
		CreatedAt:  session.CreatedAt(),
		ModifiedAt: session.ModifiedAt(),
	}
}
