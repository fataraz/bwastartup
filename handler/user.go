package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service // user->service.go
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter

	var input user.RegisterUserInput // user->input.go
	err := c.ShouldBindJSON(&input)  // untuk validasi

	// Start Error validation response
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// End Error validation response

	newUser, nil := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GenerateToken()
	formatter := user.FormatUser(newUser, "TOkenTOken")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
func (h *userHandler) Login(c *gin.Context) {
	// USER memasukan input (email & password)
	// input ditangkap handler
	// mapping dari input user ke input struct
	// input struct passing service
	// di service mencari dengan bantuan repository user dengan email x
	// mencocokan password
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedInUser, "TOkenTOken")
	response := helper.APIResponse("Successfuly logged in", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
