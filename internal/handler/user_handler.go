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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func (h *UserHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := h.service.Login(email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true) // Clear the token cookie
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
