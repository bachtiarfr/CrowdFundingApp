package main

import (
	"bwastartup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	/*
	* connection to mysql
	*/
	if err != nil {
		log.Fatal(err.Error())
	}

	/* 
	* collect the user repository
	*/
	userRepository := user.NewRepository(db)

	newUser := user.User {
		Name: "Rohim",
	}

	userRepository.SaveUser(newUser)

}