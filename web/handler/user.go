package handler

import (
	"net/http"
	"nproject/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	// c.JSON(http.StatusOK, users)
	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}
