package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/onainadapdap1/go-gin-crowfunding/auth"
	"github.com/onainadapdap1/go-gin-crowfunding/handler"
	"github.com/onainadapdap1/go-gin-crowfunding/helper"
	"github.com/onainadapdap1/go-gin-crowfunding/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:my-secret-pw-23@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// 1. memanggil repository kemudian passing db config
	userRepository := user.NewRepository(db) //return &repository{DB: db}
	// user := user.User{Name: "test simpan 1"}
	// userRepository.Save(user)
	//2. memanggil service kemudian passing repository
	userService := user.NewService(userRepository) //return &service{repo: repository}
	authService := auth.NewService()
	// fmt.Println(authService.GenerateToken(1001)) test generate token
	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNX0.njVOCJI77EykBnVyBXreHaXuaByBlUMw55rugi2UR5w")
	if err != nil {
		fmt.Println("Error")
		fmt.Println("Error")
		fmt.Println("Error")
	}
	if token.Valid {
		fmt.Println("valida")
		fmt.Println("valida")
		fmt.Println("valida")
	}
	// user := user.RegisterUserInput{Name: "test 1", Occupation: "pekerjaan 1", Email: "pegasus@gmail.com", Password: "password"}
	// userService.RegisterUser(user)
	// test upload avatar manual
	// userService.SaveAvatar(1, "images/1-profile.png")
	// 3. memanggil handler kemudian passing service sebagai parameter
	userHandler := handler.NewUserHandler(userService, authService) //	return &userHandler{userService: service}
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUserHandler)
	api.POST("/sessions", userHandler.LoginUserHandler)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	// input dari user
	// handler, mapping input dari user ke -> struct input
	// service : melakukan mapping dari struct input ke struct User
	// repository
	// db

	router.Run(":8080")
}

// ambil nilai header Authorization: Bearer tokentokentoken
// dari header Authorization: kita ambil nilai tokennya saja
// kemudian validasi token
// kita ambil nilai userID
// ambil user dari db berdasarkan userID lewat service
// kita set context isinya user

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		// ambil nilai header yang key-nya Authorization
		authHeader := c.GetHeader("Authorization")
	
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized bearer not found", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	
		// dari header Authorization kita ambil nilai tokennya saja
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		// validasi token dari authService
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized token is not valid", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		} 
		
		// ambil data user id yang ada di dalam token
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized not user is found", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		
		// ambil user dari db berdasarkan userID
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized no user found with that id", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set context isinya user yg sedang login, secara defaul mengembalikan interface
		c.Set("currentUser", user)
	}
}
