package delivery

import (
	"context"
	"fmt"
	"net/http"
	jsonRequests "profile/internal/delivery/http/user"
	"profile/internal/domain/user"
	"profile/internal/domain/user/password"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (d *Delivery) CreateUser(c *gin.Context) {
	token := c.GetHeader("x-auth-token")
	if token != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorized users can't sign-up"})
		return
	}

	request := jsonRequests.CreateUserRequest{}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := user.NewUser(request.Login, request.Password, request.Name, request.Middlename, request.Surname, request.Phone, request.City, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = d.services.User.CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, d.toResponseUser(user))
}

func (d *Delivery) UpdateUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permitted, err := checkPermissions(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !permitted {
		c.JSON(http.StatusForbidden, gin.H{"msg": "no permissions"})
		return
	}

	request := jsonRequests.UpdateUserRequest{}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upFn := func(oldUser *user.User) (*user.User, error) {
		password, err := password.EncryptPassword(request.Password)
		if err != nil {
			return nil, err
		}
		return user.NewUserWithId(oldUser.Id(), request.Login, password.String(), request.Name, request.Middlename, request.Surname, request.Phone, request.City, oldUser.Role(), oldUser.CreatedAt(), time.Now()), nil
	}

	user, err := d.services.User.UpdateUser(context.Background(), id, upFn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d.toResponseUser(user))
}

func (d *Delivery) DeleteUserById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permitted, err := checkPermissions(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !permitted {
		c.JSON(http.StatusForbidden, gin.H{"msg": "no permissions"})
		return
	}

	err = d.services.User.DeleteUserById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (d *Delivery) ReadUserById(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permitted, err := checkPermissions(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !permitted {
		c.JSON(http.StatusForbidden, gin.H{"msg": "no permissions"})
		return
	}

	user, err := d.services.User.ReadUserById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d.toResponseUser(user))
}

func (d *Delivery) ReadUserByCredetinals(c *gin.Context) {
	token := c.GetHeader("x-auth-token")
	if token != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorized users can't check creds"})
		return
	}

	request := jsonRequests.ReadUserByCredetinalsRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := d.services.User.ReadUserByCredetinals(context.Background(), request.Login, request.Pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d.toResponseUser(user))
}

func checkPermissions(c *gin.Context, userId uuid.UUID) (permitted bool, err error) {
	id, ok := c.Get("userId")
	if !ok || id == "" {
		fmt.Println("checkPermissions(): FAILED! userId not found")
		return
	}

	if userId.String() != id {
		return
	}

	permitted = true
	return
}
