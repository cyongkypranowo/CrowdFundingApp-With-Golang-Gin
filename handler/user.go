package handler

import (
	"crowdfunding/helper"
	"crowdfunding/users"
	"fmt"
	"net/http"
	"path/filepath"

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

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input users.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"error": "Server Error"}
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"isAvailable": isEmailAvailable,
	}

	metaMessage := "Email is available"

	if !isEmailAvailable {
		metaMessage = "Email has been registered"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"isUploaded": false,
		}
		response := helper.APIResponse("Avatar Upload Failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create a destination file to save the uploaded avatar
	uniqueID := helper.GenerateUniqueID()
	//get file extension
	fileExtension := filepath.Ext(file.Filename)
	//handle if fileextension is not jpg/jpeg/png will return error
	if fileExtension != ".jpg" && fileExtension != ".jpeg" && fileExtension != ".png" {
		data := gin.H{
			"isUploaded": false,
		}
		response := helper.APIResponse("Invalid File Extension", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	avatarPath := fmt.Sprintf("uploads/avatars/%s_%s.%s", uniqueID, "avatar", fileExtension)
	dst := filepath.Join("./", avatarPath)

	// Save the uploaded file to the destination
	if err := c.SaveUploadedFile(file, dst); err != nil {
		data := gin.H{
			"isUploaded": false,
		}
		response := helper.APIResponse("Failed to Save Avatar", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userID := 1
	_, err = h.userService.SaveAvatar(int64(userID), avatarPath)
	if err != nil {
		data := gin.H{
			"isUploaded": false,
		}
		response := helper.APIResponse("Failed to Save Avatar", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := gin.H{
		"isUploaded": true,
	}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
