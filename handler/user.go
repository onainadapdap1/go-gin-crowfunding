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

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	// tangkap input user
	// map input user ke struct LoginInput
	var input user.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// panggil service untuk memproses data login user
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokenlogin")
	response := helper.APIResponse("succesfully login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// objek menyimpan inputan user
	var input user.CheckEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// input dari user
	// simpan gambar di folder "images"
	// di service panggi repo
	// repo : menentukan user yng mengakses
	// repo 2 : jwt(sementara pakai hard code, yg login user id = 1)
	// repo ambil data user 1
	// repo update data user 1, simpan lokasi file
}