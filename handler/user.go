package handler

import (
	"crowdfunding/helper"
	"crowdfunding/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(user, "tokentokentoken")

	response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input users.LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := users.FormatUser(loggedInUser, "tokentokentoken")

	response := helper.APIResponse("Successfully logged in.", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
