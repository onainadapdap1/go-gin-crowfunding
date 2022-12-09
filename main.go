package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/go-gin-crowfunding/handler"
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
	// user := user.RegisterUserInput{Name: "test 1", Occupation: "pekerjaan 1", Email: "pegasus@gmail.com", Password: "password"}
	// userService.RegisterUser(user)

	// 3. memanggil handler kemudian passing service sebagai parameter
	userHandler := handler.NewUserHandler(userService) //	return &userHandler{userService: service}
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUserHandler)
	api.POST("/sessions", userHandler.LoginUserHandler)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	// input dari user
	// handler, mapping input dari user ke -> struct input
	// service : melakukan mapping dari struct input ke struct User
	// repository
	// db

	
	router.Run(":8080")
}