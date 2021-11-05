package handler

import (
	"fmt"
	"net/http"
	"nproject/user"
	"strconv"
	"time"

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

func (h *userHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	registeredUser, err := h.userService.GetUserByID(id)
	if err != nil {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Cannot get user"
		linkErr := "users"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	input := user.FormUpdateUserInput{}
	input.ID = registeredUser.ID
	input.Name = registeredUser.Name
	input.Email = registeredUser.Email
	input.Occupation = registeredUser.Occupation
	input.Role = registeredUser.Role

	c.HTML(http.StatusOK, "user_edit.html", input)

}

func (h *userHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input user.FormUpdateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "user_edit.html", input)
		return
	}

	input.ID = id

	_, err = h.userService.UpdateUser(input)
	if err != nil {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Cannot update user"
		linkErr := "users"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) NewAvatar(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	c.HTML(http.StatusOK, "user_avatar.html", gin.H{"ID": id})
}

func (h *userHandler) CreateAvatar(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	file, err := c.FormFile("avatar")
	if err != nil {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Cannot get user avatar"
		linkErr := "users"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	userID := id
	path := fmt.Sprintf("images/%d-%s-%s", userID, strconv.FormatInt(time.Now().Unix(), 10), file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Cannot upload user avatar"
		linkErr := "users"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Cannot save user avatar"
		linkErr := "users"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}
