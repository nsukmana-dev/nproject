package handler

import (
	"net/http"
	"nproject/user"
	"strconv"

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
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Error get user"
		linkErr := "users"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}
	// c.JSON(http.StatusOK, users)
	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}

func (h *userHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "user_new.html", nil)
}

func (h *userHandler) Create(c *gin.Context) {
	var input user.FormCreateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "user_new.html", input)
		return
	}

	registerInput := user.RegisterUserInput{}
	registerInput.Name = input.Name
	registerInput.Occupation = input.Occupation
	registerInput.Email = input.Email
	registerInput.Password = input.Password

	_, err = h.userService.RegisterUser(registerInput)
	if err != nil {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Error save user"
		linkErr := "users"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	c.Redirect(http.StatusFound, "/users")

}
