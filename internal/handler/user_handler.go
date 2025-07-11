package handler

import (
	"net/http"

	"simple-golang-social-media-app/internal/service"
	"simple-golang-social-media-app/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if err := h.service.Register(email, username, password); err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.Login(email, password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"error": "There was a problem loggin in. Please try again.",
		})
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"error": "Opps... Something went wrong while generating token. Please try again later.",
		})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/")
}

func (h *UserHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := h.service.Login(email, password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"error": "Invalid email or password.",
		})
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"error": "Opps... Something went wrong while generating token. Please try again later.",
		})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/")
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/login")
}
