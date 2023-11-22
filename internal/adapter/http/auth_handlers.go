package customHTTP

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project3-grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project3-grup9/internal/adapter/http/response"
	"github.com/iki-rumondor/project3-grup9/internal/application"
	"github.com/iki-rumondor/project3-grup9/internal/domain"
)

type AuthHandler struct {
	Service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var body request.Register
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	user := domain.User{
		FullName: body.FullName,
		Email:    body.Email,
		Password: body.Password,
		Role:     "member",
	}

	if err := h.Service.CreateUser(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Message{
		Message: "your account has been created successfully",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var body request.Login
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	user := domain.User{
		Email:    body.Email,
		Password: body.Password,
	}

	jwt, err := h.Service.VerifyUser(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}

func (h *AuthHandler) UpdateUser(c *gin.Context) {

	var body request.UpdateUser
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id")

	user := domain.User{
		ID:       userID,
		FullName: body.FullName,
		Email:    body.Email,
	}

	result, err := h.Service.UpdateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdateUser{
		ID:        result.ID,
		FullName:  result.FullName,
		Email:     result.Email,
		UpdatedAt: result.UpdatedAt,
	})
}

func (h *AuthHandler) DeleteUser(c *gin.Context) {

	userID := c.GetUint("user_id")

	user := domain.User{
		ID:       userID,
	}

	if err := h.Service.DeleteUser(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Your account has been successfully deleted",
	})
}


