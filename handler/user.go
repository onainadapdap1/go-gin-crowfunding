package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/go-gin-crowfunding/helper"
	"github.com/onainadapdap1/go-gin-crowfunding/user"
)

// dependensi ke service
type userHandler struct {
	userService user.Service
}

// membuat new handler
func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{userService: service}
}

func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke RegisterUserInput struct
	// struct userHandler di passing sebagai parameter service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Registered account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// memanggil service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Registered account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
