package handler

import (
	"net/http"
	"nproject/user"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *sessionHandler {
	return &sessionHandler{userService}
}

func (h *sessionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}

func (h *sessionHandler) Create(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Cannot bind user"
		linkErr := "login"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	user, err := h.userService.Login(input)
	if err != nil || user.Role != "admin" {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Cannot input user"
		linkErr := "login"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("userName", user.Name)
	session.Save()

	c.Redirect(http.StatusFound, "/users")
}

func (h *sessionHandler) Destroy(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}
