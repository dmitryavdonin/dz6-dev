package delivery

import (
	"context"
	"fmt"
	"net/http"

	jsonSession "auth/internal/delivery/http/session"
	"auth/internal/domain/session"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (d *Delivery) ReadSessionById(c *gin.Context) {

	var strId = c.Param("id")
	d.logger.Debug("Session: ReadSessionById(): id = " + strId)
	id, err := uuid.Parse(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		d.logger.Error("Session: ReadSessionById(): FAILED! Cannot parse id = " + strId + " as UUID")
		fmt.Println("Session: ReadSessionById(): FAILED! Cannot parse id = " + strId + " as UUID")
		return
	}

	session, err := d.services.Auth.ReadSessionById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		d.logger.Error("Session: ReadSessionById(): FAILED! " + err.Error())
		fmt.Println("Session: ReadSessionById(): FAILED! " + err.Error())
		return
	}

	d.logger.Debug("Session: ReadSessionById(): SUCCESS! session = " + session.Id().String())
	fmt.Println("Session: ReadSessionById(): SUCCESS! session = " + session.Id().String())

	c.JSON(http.StatusOK, d.toResponseSession(session))
}

func (d *Delivery) ReadSessionByCookie(c *gin.Context) {
	sessionCookieString, err := c.Request.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		d.logger.Error("Session: ReadSessionByCookie(): FAILED! " + err.Error())
		fmt.Println("Session: ReadSessionByCookie(): FAILED! " + err.Error())
		return
	}

	d.logger.Debug("Session: ReadSessionByCookie(): sessionCookieString = " + sessionCookieString.Value)
	fmt.Println("Session: ReadSessionByCookie(): sessionCookieString = " + sessionCookieString.Value)

	id, err := uuid.Parse(sessionCookieString.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		d.logger.Error("Session: ReadSessionByCookie(): FAILED! " + err.Error())
		fmt.Println("Session: ReadSessionByCookie(): FAILED! " + err.Error())
		return
	}

	session, err := d.services.Auth.ReadSessionById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		d.logger.Error("Session: ReadSessionByCookie(): FAILED! " + err.Error())
		fmt.Println("Session: ReadSessionByCookie(): FAILED! " + err.Error())
		return
	}

	d.logger.Debug("Session: ReadSessionByCookie(): SUCCESS! login = " + session.Login())
	d.logger.Debug("Session: ReadSessionByCookie(): SUCCESS! token = " + session.Token())

	fmt.Println("Session: ReadSessionByCookie(): SUCCESS! login = " + session.Login())
	fmt.Println("Session: ReadSessionByCookie(): SUCCESS! token = " + session.Token())

	c.Header("x-username", session.Login())
	c.Header("x-auth-token", session.Token())
	c.Header("Authorization", "Bearer "+session.Token())

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
