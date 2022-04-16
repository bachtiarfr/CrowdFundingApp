package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	
	/*
	* connection to mysql
	*/
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	/* 
	* collect the user repository
	*/
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	inputLogin := user.UserInputLogin{
		Email: "bachtiar@email.com",
		Password: "123456",
	}

	user, err := userService.LoginUser(inputLogin)
	if err != nil {
		fmt.Println("Login error")
		log.Fatal(err.Error())
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()

}