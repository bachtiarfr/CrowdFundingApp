package main

import (
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// /*
	// * connection to mysql
	// */
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Println("Connected to database")

	// /* 
	// * get user data from struck user
	// */
	// var users []user.User
	// db.Find(&users)
	
	// /* 
	// * loop through user data
	// */
	// for _, dataUser := range users {
	// 	fmt.Println(dataUser.Name)
	// }

	/* 
	* create new router
	*/
	router := gin.Default()
	router.GET("/", controller)
	router.Run()
}

/* 
* Routing controller
*/
func controller(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	/*
	* connection to mysql
	*/
	if err != nil {
		log.Fatal(err.Error())
	}

	/* 
	* get user data from struck user
	*/
	var users []user.User
	db.Find(&users)
	
	c.JSON(200, users)
}